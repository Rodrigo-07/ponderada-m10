package service

import "ponderada1/internal/client"

type WeatherService struct {
	om *client.OpenMeteoClient
}

func NewWeatherService() *WeatherService {
	return &WeatherService{om: client.NewOpenMeteoClient()}
}

// Retorna temperatura atual simplificada
func (s *WeatherService) GetCurrentTemp(lat, lon float64) (float64, error) {
	return s.om.CurrentTemp(lat, lon)
}
