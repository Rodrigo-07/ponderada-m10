package client

import (
	"fmt"
	"ponderada1/internal/model"

	"github.com/go-resty/resty/v2"
)

type OpenDeckClient struct {
	baseURL string
	rest    *resty.Client
}

func NewOpenDeckClient() *OpenDeckClient {
	return &OpenDeckClient{
		baseURL: "https://deckofcardsapi.com/api/deck/",
		rest:    resty.New(),
	}
}

func (c *OpenDeckClient) ShuffleNewDeck(deckCount int) (*model.Deck, error) {
	url := fmt.Sprintf("%snew/shuffle/?deck_count=%d", c.baseURL, deckCount)

	var deck model.Deck
	resp, err := c.rest.R().
		SetResult(&deck).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("erro na requisição: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("erro da API: %s", resp.Status())
	}

	return &deck, nil
}

func (c *OpenDeckClient) DrawCards(deckID string, count int) (*model.DrawnCards, error) {
	url := fmt.Sprintf("%s%s/draw/?count=%d", c.baseURL, deckID, count)

	fmt.Printf("Fazendo requisição para:")
	fmt.Println(url)

	var drawnCards model.DrawnCards
	resp, err := c.rest.R().
		SetResult(&drawnCards).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("erro na requisição: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("erro da API: %s", resp.Status())
	}

	return &drawnCards, nil
}
