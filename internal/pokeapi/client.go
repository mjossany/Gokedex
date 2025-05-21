package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type PokeApiClient struct {
	httpClient http.Client
}

func NewPokeApiClient(timeout time.Duration) *PokeApiClient {
	return &PokeApiClient{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

type PokeApi interface {
	ListLocations(pageURL *string) (RespShallowLocations, error)
}

func (c *PokeApiClient) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationResp, nil
}
