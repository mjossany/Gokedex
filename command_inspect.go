package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemonName *string) error {
	cfg.nextLocationsURL = nil
	pokemonInfo, found := cfg.pokedex.Pokedex[*pokemonName]
	if !found {
		fmt.Println("You have not caught a", *pokemonName)
		return nil
	}

	fmt.Println("Name: ", pokemonInfo.Name)
	fmt.Println("Height: ", pokemonInfo.Height)
	fmt.Println("Weight: ", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemonInfo.Types {
		fmt.Printf("  - %s\n", typ.Type.Name)
	}

	return nil
}
