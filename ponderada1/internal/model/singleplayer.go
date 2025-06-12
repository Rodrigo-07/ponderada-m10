package model

import "github.com/lib/pq"

type StringArray []string

type Singleplayer struct {
	GameID     *string        `gorm:"primaryKey;default:gen_random_uuid()" json:"game_id,omitempty"`
	PlayerName string         `json:"player_name"`
	Result     string         `json:"result"`
	CardsSum   int            `json:"card_sum"`
	DeckID     string         `json:"deck_id"`
	DrawnCards pq.StringArray `gorm:"type:text[]" json:"drawn_cards" swaggertype:"array,string"`
	Moves      pq.StringArray `gorm:"type:text[]" json:"moves" swaggertype:"array,string"`
}

type CreateSinlgePlayerGameRequest struct {
	PlayerName string `json:"player_name" binding:"required"`
}

type MakeMoveSinglePlayerRequest struct {
	GameID string `json:"game_id" binding:"required"`
	Move   string `json:"move" binding:"required"`
}
