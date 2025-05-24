package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, pokemonName *string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", *pokemonName)
	cfg.nextLocationsURL = nil
	pokemonInfoResp, err := cfg.pokeapiClient.FetchPokemonInfo(cfg.nextLocationsURL,
		*pokemonName)
	if err != nil {
		return err
	}

	const catchThreshold = 50

	chance := rand.Intn(pokemonInfoResp.BaseExperience)

	if chance < catchThreshold {
		fmt.Printf("%s was caught!\n", pokemonInfoResp.Name)
		cfg.pokedex.Pokedex[pokemonInfoResp.Name] = pokemonInfoResp
	} else {
		fmt.Printf("%s escaped!\n", pokemonInfoResp.Name)
	}

	return nil
}
