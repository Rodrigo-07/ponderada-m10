package model

// import (
// 	"github.com/lib/pq"
// )

type StringArray []string

type Singleplayer struct {
	GameID     string      `json:"game_id" gorm:"primaryKey"`
	PlayerName string      `json:"player_name"`
	Result     string      `json:"result"`
	DeckID     string      `json:"deck_id"`
	DrawnCards StringArray `json:"drawn_cards" gorm:"type:text[]"`
	Moves      StringArray `json:"moves" gorm:"type:text[]"`
}