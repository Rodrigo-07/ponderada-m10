package repository

import (
	"ponderada1/internal/db"
	"ponderada1/internal/model"
)

func InsertSingleplayerGame(game *model.Singleplayer) error {
	return db.GetDB().Create(game).Error
}

func ListSingleplayerGames() ([]model.Singleplayer, error) {
	var games []model.Singleplayer
	err := db.GetDB().Find(&games).Error
	return games, err
}
