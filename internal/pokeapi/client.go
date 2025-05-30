package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/mjossany/Gokedex/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type PokeApiClient struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewPokeApiClient(timeout time.Duration, cache pokecache.Cache) *PokeApiClient {
	return &PokeApiClient{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}
}

type PokeApi interface {
	ListLocations(pageURL *string) (RespShallowLocations, error)
	ListLocationPokemonEncounters(pageURL *string, locationName string) (RespLocationPokemonEncounters, error)
	FetchPokemonInfo(pageURL *string, pokemonName string) (RespPokemonInfo, error)
}

func (c *PokeApiClient) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cachedData, found := c.cache.Get(url)
	if found {
		locationResp := RespShallowLocations{}
		err := json.Unmarshal(cachedData, &locationResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationResp, nil
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

	c.cache.Add(url, data)

	locationResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationResp, nil
}

func (c *PokeApiClient) ListLocationPokemonEncounters(pageURL *string, locationName string) (RespLocationPokemonEncounters, error) {
	url := baseURL + "/location-area/" + locationName
	if pageURL != nil {
		url = *pageURL
	}

	cachedData, found := c.cache.Get(url)
	if found {
		locationPokemonEncounters := RespLocationPokemonEncounters{}
		err := json.Unmarshal(cachedData, &locationPokemonEncounters)
		if err != nil {
			return RespLocationPokemonEncounters{}, err
		}
		return locationPokemonEncounters, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationPokemonEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationPokemonEncounters{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationPokemonEncounters{}, err
	}

	c.cache.Add(url, data)

	locationPokemonEncounters := RespLocationPokemonEncounters{}
	err = json.Unmarshal(data, &locationPokemonEncounters)
	if err != nil {
		return RespLocationPokemonEncounters{}, err
	}

	return locationPokemonEncounters, nil
}

func (c *PokeApiClient) FetchPokemonInfo(pageURL *string, pokemonName string) (RespPokemonInfo, error) {
	url := baseURL + "/pokemon/" + pokemonName
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonInfo{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	pokemonInfo := RespPokemonInfo{}
	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	return pokemonInfo, nil
}
