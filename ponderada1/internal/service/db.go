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

func (s *GameService) CreateSingleplayer(game model.Singleplayer) (*model.Singleplayer, error) {
	createdGame, err := repository.InsertSingleplayerGame(&game)
	if err != nil {
		return nil, err
	}
	return createdGame, nil
}

func (s *GameService) GetAllSingleplayerGames() ([]model.Singleplayer, error) {
	return repository.ListSingleplayerGames()
}

func (s *GameService) GetSingleplayerGameByID(gameID string) (*model.Singleplayer, error) {
	return repository.GetSingleplayerGameByID(gameID)
}

func (s *GameService) UpdateSingleplayerDrawnCards(gameID string, drawnCards []string, CardsSum int) (*model.Singleplayer, error) {
	var game *model.Singleplayer
	game, err := repository.UpdateSingleplayerDrawnCards(gameID, drawnCards, CardsSum)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (s *GameService) UpdateSingleplayerResult(gameID string, result string) (*model.Singleplayer, error) {
	var game *model.Singleplayer
	game, err := repository.UpdateSingleplayerResult(gameID, result)
	if err != nil {
		return nil, err
	}
	return game, nil
}
