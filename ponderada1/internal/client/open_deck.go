package client

import (
	"fmt"

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

func (c *OpenDeckClient) ShuffleNewDeck(deckCount int) (string, error) {
	url := fmt.Sprintf("%snew/shuffle/?deck_count=%d", c.baseURL, deckCount)
	resp, err := c.rest.R().Get(url)
	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", fmt.Errorf("error response: %s", resp.Status())
	}

	return resp.String(), nil
}
