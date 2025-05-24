package main

import (
	"time"

	"github.com/mjossany/Gokedex/internal/pokeapi"
	"github.com/mjossany/Gokedex/internal/pokecache"
	"github.com/mjossany/Gokedex/internal/pokedex"
)

func main() {
	cacheInterval := 5 * time.Minute
	appCache := pokecache.NewCache(cacheInterval)

	pokeApi := pokeapi.NewPokeApiClient(5*time.Second, *appCache)
	pokedex := pokedex.NewPokedex()
	cfg := &config{
		pokeapiClient: pokeApi,
		pokedex:       pokedex,
	}
	startRepl(cfg)
}
