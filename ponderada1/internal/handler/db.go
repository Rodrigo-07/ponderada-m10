// internal/handler/game_handler.go
package handler

import (
	"net/http"
	"ponderada1/internal/model"
	"ponderada1/internal/service"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	svc *service.GameService
}

func RegisterGameRoutes(rg *gin.RouterGroup) {
	handler := &GameHandler{
		svc: service.NewGameService(),
	}

	rg.GET("/get-games", handler.listGames)
	rg.POST("/create-game", handler.createGame)
}

// @Summary Lista todos os jogos singleplayer
// @Description Retorna todos os registros de jogos
// @Tags games
// @Produce json
// @Success 200 {array} model.Singleplayer
// @Failure 500 {object} map[string]string
// @Router /get-games [get]
func (h *GameHandler) listGames(c *gin.Context) {
	games, err := h.svc.GetAllSingleplayerGames()
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
// @Param game body model.Singleplayer true "Dados do jogo"
// @Success 201 {object} model.Singleplayer
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /create-game [post]
func (h *GameHandler) createGame(c *gin.Context) {
	var game model.Singleplayer
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}
	if err := h.svc.CreateSingleplayer(game); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao salvar jogo"})
		return
	}
	c.JSON(http.StatusCreated, game)
}
