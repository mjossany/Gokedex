package main

import (
	"fmt"
)

func commandPokedex(cfg *config, pokemonName *string) error {
	cfg.nextLocationsURL = nil
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex.Pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
