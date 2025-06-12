package repository

import (
	"ponderada1/internal/db"
	"ponderada1/internal/model"

	"github.com/lib/pq"
)

func InsertSingleplayerGame(game *model.Singleplayer) (*model.Singleplayer, error) {
	err := db.GetDB().Create(game).Error
	if err != nil {
		return nil, err
	}
	return game, nil
}

func ListSingleplayerGames() ([]model.Singleplayer, error) {
	var games []model.Singleplayer
	err := db.GetDB().Find(&games).Error
	return games, err
}

func UpdateSingleplayerDrawnCards(gameID string, newCards []string, cardsSum int) (*model.Singleplayer, error) {
	var game model.Singleplayer

	// Fetch the current game and its saved cards
	err := db.GetDB().Where("game_id = ?", gameID).First(&game).Error
	if err != nil {
		return nil, err
	}

	// Append the new cards to the existing ones
	game.DrawnCards = append(game.DrawnCards, newCards...)
	game.CardsSum = cardsSum

	// Update the fields in the database
	err = db.GetDB().Model(&game).Where("game_id = ?", gameID).
		Updates(map[string]interface{}{
			"drawn_cards": pq.StringArray(game.DrawnCards),
			"cards_sum":   game.CardsSum,
		}).Error
	if err != nil {
		return nil, err
	}

	err = db.GetDB().Where("game_id = ?", gameID).First(&game).Error
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func UpdateSingleplayerResult(gameID string, result string) (*model.Singleplayer, error) {
	var game model.Singleplayer
	err := db.GetDB().Model(&game).Where("game_id = ?", gameID).Update("result", result).Error
	if err != nil {
		return nil, err
	}

	err = db.GetDB().Where("game_id = ?", gameID).First(&game).Error
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func GetSingleplayerGameByID(gameID string) (*model.Singleplayer, error) {
	var game model.Singleplayer
	err := db.GetDB().Where("game_id = ?", gameID).First(&game).Error
	if err != nil {
		return nil, err
	}
	return &game, nil
}
