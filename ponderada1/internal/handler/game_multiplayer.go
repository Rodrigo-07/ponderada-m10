package handler

import (
	"net/http"

	"ponderada1/internal/model"
	"ponderada1/internal/service"

	"github.com/gin-gonic/gin"
)

type MultiplayerGameHandler struct {
	gameSvc *service.GameService
	deckSvc *service.DeckService
}

func RegisterMultiplayerGameRoutes(rg *gin.RouterGroup) {
	h := &MultiplayerGameHandler{
		gameSvc: service.NewGameService(),
		deckSvc: service.NewDeckService(),
	}

	rg.GET("/get-multiplayer-games", h.listGamesMulti)
	rg.POST("/create-game-multiplayer", h.createGameMultiplayer)
	rg.POST("/make-move-multiplayer", h.makeMoveMultiplayer)
}

// @Summary Lista todos os jogos multiplayer
// @Description Retorna todos os registros de jogos multiplayer
// @Tags games
// @Produce json
// @Success 200 {array} model.Multiplayer
// @Failure 500 {object} map[string]string
// @Router /get-multiplayer-games [get]
func (h *MultiplayerGameHandler) listGamesMulti(c *gin.Context) {
	games, err := h.gameSvc.GetAllMultiplayerGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar jogos"})
		return
	}
	c.JSON(http.StatusOK, games)
}

func (h *MultiplayerGameHandler) getMultiplayerGame(c *gin.Context) {
	id := c.Param("id")
	requester := c.Query("player") // opcional; se vazio, esconde ambas

	game, err := h.gameSvc.GetMultiplayerGameByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "jogo não encontrado"})
		return
	}

	c.JSON(200, sanitizeForPlayer(*game, requester))
}

// @Summary Cria um novo jogo multiplayer
// @Description Inicializa um jogo multiplayer com dois jogadores, cada um utilizando um baralho independente.
// @Tags games
// @Accept json
// @Produce json
// @Param body body model.CreateMultiplayerGameRequest true "Dados para criação do jogo multiplayer"
// @Success 201 {object} model.Multiplayer
// @Failure 400 {object} map[string]string "dados inválidos"
// @Failure 500 {object} map[string]string "erro interno do servidor"
// @Router /create-game-multiplayer [post]
func (h *MultiplayerGameHandler) createGameMultiplayer(c *gin.Context) {
    var req model.CreateMultiplayerGameRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
        return
    }

    deck, err := h.deckSvc.ShuffleNewDeck(1)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    /* ----- compra exatamente 2 cartas para cada jogador ----- */
    p1, err := h.deckSvc.DrawCards(deck.DeckID, 2)
    if err != nil || len(p1.Cards) < 2 {
        c.JSON(500, gin.H{"error": "falha ao comprar cartas p/ jogador 1"})
        return
    }
    p2, err := h.deckSvc.DrawCards(deck.DeckID, 2)
    if err != nil || len(p2.Cards) < 2 {
        c.JSON(500, gin.H{"error": "falha ao comprar cartas p/ jogador 2"})
        return
    }

    vis1 := model.StringArray{p1.Cards[0].Code}
    hid1 := p1.Cards[1].Code

    vis2 := model.StringArray{p2.Cards[0].Code}
    hid2 := p2.Cards[1].Code

    s1 := h.deckSvc.CardValue(p1.Cards[0].Code) +
          h.deckSvc.CardValue(p1.Cards[1].Code)

    s2 := h.deckSvc.CardValue(p2.Cards[0].Code) +
          h.deckSvc.CardValue(p2.Cards[1].Code)

    game := model.Multiplayer{
        Player1Name: req.Player1Name,
        Player2Name: req.Player2Name,
        Result:      "in_progress",
        DeckID:      deck.DeckID,

        Player1VisibleCards: vis1, Player1HiddenCard: hid1,
        Player2VisibleCards: vis2, Player2HiddenCard: hid2,

        Player1ExtraCards: model.StringArray{}, Player2ExtraCards: model.StringArray{},
        Player1Moves: model.StringArray{},     Player2Moves:      model.StringArray{},
        Player1Score: s1, Player2Score: s2,

        CurrentTurn: "player1",
        Round:       1,
    }

    created, err := h.gameSvc.CreateMultiplayer(game)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, created)
}

// @Summary Realiza uma jogada em um jogo multiplayer
// @Description Permite que um jogador realize uma jogada em um jogo multiplayer. As jogadas podem ser "draw", "pass" ou "stop".
// @Tags games
// @Accept json
// @Produce json
// @Param body body model.MakeMoveMultiplayerRequest true "Dados para realizar a jogada"
// @Success 200 {object} model.Multiplayer
// @Failure 400 {object} map[string]string "dados inválidos ou erro de jogada"
// @Failure 500 {object} map[string]string "erro interno do servidor"
// @Router /make-move-multiplayer [post]
func (h *MultiplayerGameHandler) makeMoveMultiplayer(c *gin.Context) {
	/* ---------- 1. Bind & fetch ---------- */
	var req model.MakeMoveMultiplayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	game, err := h.gameSvc.GetMultiplayerGameByID(req.GameID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if game.Result != "in_progress" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jogo já finalizado"})
		return
	}

	/* ---------- 2. Identifica jogador ---------- */
	var (
		isP1   bool
		moves  *model.StringArray
		extras *model.StringArray
		score  *int
		stop   *bool
	)
	switch req.PlayerName {
	case game.Player1Name:
		isP1 = true
		moves, extras, score, stop =
			&game.Player1Moves, &game.Player1ExtraCards, &game.Player1Score, &game.Player1Stop
	case game.Player2Name:
		moves, extras, score, stop =
			&game.Player2Moves, &game.Player2ExtraCards, &game.Player2Score, &game.Player2Stop
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "nome de jogador inválido"})
		return
	}

	/* ---------- 2.1 Confere vez ---------- */
	if (isP1 && game.CurrentTurn != "player1") ||
		(!isP1 && game.CurrentTurn != "player2") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "não é sua vez"})
		return
	}

	/* ---------- 3. Máx. 3 jogadas ---------- */
	if len(*moves) >= 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "máx 3 rodadas"})
		return
	}

	/* ---------- 4. Executa ação ---------- */
	switch req.Move {
	case "draw":
		draw, derr := h.deckSvc.DrawCards(game.DeckID, 1)
		if derr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": derr.Error()})
			return
		}
		card := draw.Cards[0]

		*extras = append(*extras, card.Code)
		*score += h.deckSvc.CardValue(card.Code)
		*moves = append(*moves, "draw")

		// estourou  -> define vencedor imediatamente
		if *score > 21 {
			*stop = true
		}

	case "pass":
		*moves = append(*moves, "pass")

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "movimento inválido"})
		return
	}

	/* ---------- 5. Avança turno/rodada se jogo continua ---------- */
	if game.Result == "in_progress" {
		if game.CurrentTurn == "player1" {
			game.CurrentTurn = "player2"
		} else {
			game.CurrentTurn = "player1"
			game.Round++ // nova rodada quando os dois já jogaram
		}

		// Finaliza se chegaram a 3 rodadas ou ambos pararam manualmente
		if game.Round > 3 || (game.Player1Stop && game.Player2Stop) {
			game.Result = decideWinner(game)
		}
	}

	/* ---------- 6. Persiste ---------- */
	updated, err := h.gameSvc.UpdateMultiplayer(req.GameID, *game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func codesAndScore(cards []model.Card, deckSvc *service.DeckService) ([]string, int) {
	var codes []string
	sum := 0
	for _, c := range cards {
		codes = append(codes, c.Code)
		sum += deckSvc.CardValue(c.Code)
	}
	return codes, sum
}

func decideWinner(g *model.Multiplayer) string {
	s1, s2 := g.Player1Score, g.Player2Score

	best := func(a, b int) int {
		if a > 21 && b > 21 {
			if a < b {
				return a
			}
			if b < a {
				return b
			}
			return a
		}
		if a > 21 {
			return b
		}
		if b > 21 {
			return a
		}
		if 21-a < 21-b {
			return a
		}
		if 21-b < 21-a {
			return b
		}
		return a
	}

	winnerScore := best(s1, s2)
	switch {
	case s1 == s2:
		return "draw"
	case winnerScore == s1:
		return "player1"
	default:
		return "player2"
	}
}

func sanitizeForPlayer(g model.Multiplayer, requester string) model.Multiplayer {
	clone := g
	if requester == g.Player1Name {
		clone.Player2HiddenCard = ""
	} else if requester == g.Player2Name {
		clone.Player1HiddenCard = ""
	} else {
		clone.Player1HiddenCard = ""
		clone.Player2HiddenCard = ""
	}
	return clone
}
