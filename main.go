package main

import (
	"time"

	"github.com/mjossany/Gokedex/internal/pokeapi"
	"github.com/mjossany/Gokedex/internal/pokecache"
)

func main() {
	cacheInterval := 5 * time.Minute
	appCache := pokecache.NewCache(cacheInterval)

	pokeApi := pokeapi.NewPokeApiClient(5*time.Second, *appCache)
	cfg := &config{
		pokeapiClient: pokeApi,
	}
	startRepl(cfg)
}
