package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Endpoint simples para previsão do tempo: https://open-meteo.com/
type OpenMeteoClient struct {
	baseURL string
	rest    *resty.Client
}

func NewOpenMeteoClient() *OpenMeteoClient {
	return &OpenMeteoClient{
		baseURL: "https://api.open-meteo.com/v1",
		rest:    resty.New(),
	}
}

func (c *OpenMeteoClient) CurrentTemp(lat, lon float64) (float64, error) {
	var resp struct {
		Current struct {
			Temperature float64 `json:"temperature_2m"`
		} `json:"current"`
	}

	_, err := c.rest.R().
		SetQueryParams(map[string]string{
			"latitude":  fmt.Sprint(lat),
			"longitude": fmt.Sprint(lon),
			"current":   "temperature_2m",
		}).
		SetResult(&resp).
		Get(c.baseURL + "/forecast")
	if err != nil {
		return 0, err
	}
	return resp.Current.Temperature, nil
}
