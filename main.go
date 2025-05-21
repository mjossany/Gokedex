package main

import (
	"time"

	"github.com/mjossany/Gokedex/internal/pokeapi"
)

func main() {
	pokeApi := pokeapi.NewPokeApiClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeApi,
	}
	startRepl(cfg)
}
