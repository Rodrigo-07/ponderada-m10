package handler

import (
	"net/http"
	"strconv"

	"ponderada1/internal/service"

	"github.com/gin-gonic/gin"
)

type DeckHandler struct {
	svc *service.DeckService
}

func RegisterDeckRoutes(rg *gin.RouterGroup) {
	h := DeckHandler{svc: service.NewDeckService()}
	rg.GET("/shuffle", h.shuffle)
}

// @Summary      Embaralhar baralho
// @Description  Embaralha um novo baralho de cartas
// @Tags         deck
// @Param        count  query  int  false  "Número de baralhos (default: 1)"
// @Success      200  {string}  string  "Baralho embaralhado"
// @Failure      400  {object}  map[string]string
// @Router       /deck/shuffle [get]
func (h *DeckHandler) shuffle(c *gin.Context) {
	countStr := c.Query("count")
	count := 1 // valor padrão
	if countStr != "" {
		var err error
		count, err = strconv.Atoi(countStr)
		if err != nil || count <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "count inválido"})
			return
		}
	}

	result, err := h.svc.ShuffleNewDeck(count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, result.DeckID)
}
