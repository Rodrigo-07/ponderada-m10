package model

type Multiplayer struct {
	GameID     string   `gorm:"primaryKey;default:gen_random_uuid()" json:"game_id"`
	Player1Name       string   `json:"player1_name"`
	Player2Name       string   `json:"player2_name"`
	Result            string   `json:"result"` // "player1", "player2", or "draw"
	Player1DeckID     string   `json:"player1_deck_id"`
	Player2DeckID     string   `json:"player2_deck_id"`
	Player1DrawnCards []string `json:"player1_drawn_cards"`
	Player2DrawnCards []string `json:"player2_drawn_cards"`
	Player1Moves      []string `json:"player1_moves"`
	Player2Moves      []string `json:"player2_moves"`
}
