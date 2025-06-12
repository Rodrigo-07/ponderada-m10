package service

import (
	"fmt"
	"ponderada1/internal/client"
	"ponderada1/internal/model"
	"strconv"
)

type DeckService struct {
	om *client.OpenDeckClient
}

func NewDeckService() *DeckService {
	return &DeckService{
		om: client.NewOpenDeckClient(),
	}
}

func (s *DeckService) ShuffleNewDeck(deckCount int) (*model.Deck, error) {
	deck, err := s.om.ShuffleNewDeck(deckCount)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Baralho embaralhado com ID: %s\n", deck.DeckID)
	return deck, nil
}

func (s *DeckService) DrawCards(deckID string, count int) (*model.DrawnCards, error) {
	drawnCards, err := s.om.DrawCards(deckID, count)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Cartas do baralho %s: %d cartas restantes\n", drawnCards.DeckID, drawnCards.Remaining)
	return drawnCards, nil
}

func (s *DeckService) CardValue(code string) int {
	valuePart := code[:1]

	fmt.Printf("Calculando valor da carta: %s\n", valuePart)

	switch valuePart {
	case "0":
		return 10
	case "A":
		return 1
	case "K", "Q", "J":
		return 10
	default:
		if n, err := strconv.Atoi(valuePart); err == nil {
			return n
		}
		return 0
	}
}
