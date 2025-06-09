package model

type Deck struct {
	Success   bool   `json:"success"`
	DeckID    string `json:"deck_id"`
	Remaining int    `json:"remaining"`
	Shuffled  bool   `json:"shuffled"`
}
