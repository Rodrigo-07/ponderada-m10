package model

import "github.com/lib/pq"

type StringArray = pq.StringArray
type Multiplayer struct {
    GameID string `gorm:"primaryKey;default:gen_random_uuid()" json:"game_id"`

    Player1Name string `json:"player1_name"`
    Player2Name string `json:"player2_name"`
    Result      string `json:"result"`           // in_progress | player1 | player2 | draw

    DeckID      string `json:"deck_id"`
    CurrentTurn string `json:"current_turn"`     // player1 | player2
    Round       int    `json:"round"`            // 1-based (1,2,3)

	Player1VisibleCards pq.StringArray `gorm:"type:text[]" json:"player1_visible" swaggertype:"array,string"`
    Player1HiddenCard   string         `json:"player1_hidden"`
    Player2VisibleCards pq.StringArray `gorm:"type:text[]" json:"player2_visible" swaggertype:"array,string"`
    Player2HiddenCard   string         `json:"player2_hidden"`

    Player1ExtraCards pq.StringArray `gorm:"type:text[]" json:"player1_extra,omitempty" swaggertype:"array,string"`
    Player2ExtraCards pq.StringArray `gorm:"type:text[]" json:"player2_extra,omitempty" swaggertype:"array,string"`

    Player1Moves pq.StringArray `gorm:"type:text[]" json:"player1_moves" swaggertype:"array,string"`
    Player2Moves pq.StringArray `gorm:"type:text[]" json:"player2_moves" swaggertype:"array,string"`

    Player1Score int  `json:"player1_score"`
    Player2Score int  `json:"player2_score"`
    Player1Stop  bool `json:"player1_stop"`
    Player2Stop  bool `json:"player2_stop"`
}


type CreateMultiplayerGameRequest struct {
	Player1Name string `json:"player1_name" binding:"required"`
	Player2Name string `json:"player2_name" binding:"required"`
}

type MakeMoveMultiplayerRequest struct {
	GameID     string `json:"game_id"     binding:"required,uuid"`
	PlayerName string `json:"player_name" binding:"required"`
	Move       string `json:"move"        binding:"required,oneof=draw pass stop"`
}
