// internal/handler/game_handler.go
package handler

import (
	"fmt"
	"net/http"
	"ponderada1/internal/model"
	"ponderada1/internal/service"

	"github.com/gin-gonic/gin"
)

type SinglePlayerGameHandler struct {
	gameSvc *service.GameService
	deckSvc *service.DeckService
}

func RegisterSinglePlayerGameRoutes(rg *gin.RouterGroup) {
	handler := &SinglePlayerGameHandler{
		gameSvc: service.NewGameService(),
		deckSvc: service.NewDeckService(),
	}

	rg.GET("/get-games", handler.listGames)
	rg.POST("/create-game", handler.createGameSinglePlayer)
	rg.POST("/make-move-singleplayer", handler.makeMoveSinglePlayer)
}

// @Summary Lista todos os jogos singleplayer
// @Description Retorna todos os registros de jogos
// @Tags games
// @Produce json
// @Success 200 {array} model.Singleplayer
// @Failure 500 {object} map[string]string
// @Router /get-games [get]
func (h *SinglePlayerGameHandler) listGames(c *gin.Context) {
	games, err := h.gameSvc.GetAllSingleplayerGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar jogos"})
		return
	}
	c.JSON(http.StatusOK, games)
}

// @Summary Cria um novo jogo singleplayer
// @Description Cria e persiste um novo jogo no banco
// @Tags games
// @Accept json
// @Produce json
// @Param game body model.CreateSinlgePlayerGameRequest true "Dados do jogo"
// @Success 201 {object} model.Singleplayer
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /create-game [post]
func (h *SinglePlayerGameHandler) createGameSinglePlayer(c *gin.Context) {
	var req model.CreateSinlgePlayerGameRequest

	// Faz o bind do JSON para o DTO correto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	deckResp, err := h.deckSvc.ShuffleNewDeck(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao embaralhar baralho"})
		return
	}

	// Give 3 cards to the player
	drawnCards, err := h.deckSvc.DrawCards(deckResp.DeckID, 2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao desenhar cartas"})
		return
	}
	var codes []string
	for _, card := range drawnCards.Cards {
		codes = append(codes, card.Code)
	}

	// Get the values of the drawn cards and calculate the sum
	cardValuesSum := 0
	for _, card := range drawnCards.Cards {
		fmt.Printf("Carta desenhada: %s, Valor: %d\n", card.Code, h.deckSvc.CardValue(card.Code))
		cardValuesSum += h.deckSvc.CardValue(card.Code)
	}

	var result string

	if cardValuesSum > 21 {
		result = "busted"
	} else if cardValuesSum == 21 {
		result = "won"
	} else {
		result = "in_progress"
	}

	game := model.Singleplayer{
		PlayerName: req.PlayerName,
		Result:     result,
		DeckID:     deckResp.DeckID,
		DrawnCards: codes,
		CardsSum:   cardValuesSum,
	}

	createdGame, err := h.gameSvc.CreateSingleplayer(game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar jogo"})
		return
	}
	c.JSON(http.StatusCreated, createdGame)
}

// @Summary Cria um novo jogo singleplayer
// @Description Cria e persiste um novo jogo no banco
// @Tags games
// @Accept json
// @Produce json
// @Param game body model.MakeMoveSinglePlayerRequest true "Dados do jogo"
// @Success 201 {object} model.Singleplayer
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /make-move-singleplayer [post]
func (h *SinglePlayerGameHandler) makeMoveSinglePlayer(c *gin.Context) {
	var req model.MakeMoveSinglePlayerRequest

	// Faz o bind do JSON para o DTO correto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	if req.Move == "draw" {
		// Lógica para desenhar cartas
		gameinfo, err := h.gameSvc.GetSingleplayerGameByID(req.GameID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar jogo"})
			return
		}

		if gameinfo.Result != "in_progress" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "jogo já finalizado"})
			return
		}

		drawnCards, err := h.deckSvc.DrawCards(*&gameinfo.DeckID, 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var codes []string
		for _, card := range drawnCards.Cards {
			codes = append(codes, card.Code)
		}

		// Get card values and calculate the sum
		cardValuesSum := 0
		for _, card := range drawnCards.Cards {
			cardValuesSum += h.deckSvc.CardValue(card.Value)
		}

		cardValuesSum += gameinfo.CardsSum

		game, err := h.gameSvc.UpdateSingleplayerDrawnCards(req.GameID, codes, cardValuesSum)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao atualizar jogo com cartas desenhadas"})
			return
		}

		if cardValuesSum > 21 {
			game, err := h.gameSvc.UpdateSingleplayerResult(req.GameID, "busted")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao atualizar jogo com resultado"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Você estourou! Jogo terminado.", "game": game})
			return
		} else if cardValuesSum == 21 {
			game, err := h.gameSvc.UpdateSingleplayerResult(req.GameID, "won")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao atualizar jogo com resultado"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Você venceu! Jogo terminado.", "game": game})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "O jogo continua", "drawnCards": drawnCards, "game": game})
		return
	}
	if req.Move == "stop" {
		// Lógica para parar o jogo
		game, err := h.gameSvc.UpdateSingleplayerResult(req.GameID, "stopped")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, game)
		return
	}

}
