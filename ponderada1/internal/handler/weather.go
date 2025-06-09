package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"ponderada1/internal/service"
)

type WeatherHandler struct {
	svc *service.WeatherService
}

func RegisterWeatherRoutes(rg *gin.RouterGroup) {
	h := WeatherHandler{svc: service.NewWeatherService()}
	rg.GET("/weather/current", h.current)
}

// @Summary      Temperatura atual
// @Description  Consulta a temperatura atual via Open-Meteo
// @Tags         weather
// @Param        lat  query  number  true  "Latitude"
// @Param        lon  query  number  true  "Longitude"
// @Success      200  {object}  map[string]float64
// @Failure      400  {object}  map[string]string
// @Router       /weather/current [get]
func (h *WeatherHandler) current(c *gin.Context) {
	lat, err1 := strconv.ParseFloat(c.Query("lat"), 64)
	lon, err2 := strconv.ParseFloat(c.Query("lon"), 64)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lat/lon inválidos"})
		return
	}

	temp, err := h.svc.GetCurrentTemp(lat, lon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"temperature": temp})
}
