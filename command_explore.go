package main

import (
	"fmt"
)

func commandExplore(cfg *config, secondArgument *string) error {
	fmt.Printf("Exploring %s...\n", *secondArgument)
	cfg.nextLocationsURL = nil
	locationPokemonEncounterResp, err := cfg.pokeapiClient.ListLocationPokemonEncounters(cfg.nextLocationsURL,
		*secondArgument)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range locationPokemonEncounterResp.PokemonEncounters {
		fmt.Println("- ", pokemonEncounter.Pokemon.Name)
	}
	return nil
}
