package model

type Deck struct {
	Success   bool   `json:"success"`
	DeckID    string `json:"deck_id"`
	Remaining int    `json:"remaining"`
	Shuffled  bool   `json:"shuffled"`
}

type Card struct {
	Code  string `json:"code"`
	Image string `json:"image"`
	Images struct {
		SVG string `json:"svg"`
		PNG string `json:"png"`
	} `json:"images"`
	Value string `json:"value"`
	Suit  string `json:"suit"`
}

type DrawnCards struct {
	Success   bool   `json:"success"`
	DeckID    string `json:"deck_id"`
	Cards     []Card `json:"cards"`
	Remaining int    `json:"remaining"`
}
