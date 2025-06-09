package service

import "ponderada1/internal/client"

type DeckService struct {
	om *client.OpenDeckClient
}

func NewDeckService() *DeckService {
	return &DeckService{om: client.NewOpenDeckClient()}
}

// Embaralha um novo baralho de cartas
func (s *DeckService) ShuffleNewDeck(deckCount int) (string, error) {
	return s.om.ShuffleNewDeck(deckCount)
}
