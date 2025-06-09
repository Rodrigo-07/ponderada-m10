// internal/service/singleplayer.go
package service

import (
	"ponderada1/internal/model"
	"ponderada1/internal/repository"
)

type GameService struct{}

func NewGameService() *GameService {
	return &GameService{}
}

func (s *GameService) CreateSingleplayer(game model.Singleplayer) error {
	return repository.InsertSingleplayerGame(&game)
}

func (s *GameService) GetAllSingleplayerGames() ([]model.Singleplayer, error) {
	return repository.ListSingleplayerGames()
}
