package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mjossany/Gokedex/internal/pokeapi"
	"github.com/mjossany/Gokedex/internal/pokedex"
)

type config struct {
	pokeapiClient        pokeapi.PokeApi
	nextLocationsURL     *string
	previousLocationsURL *string
	pokedex              *pokedex.Pokedex
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		var secondArgument *string
		if len(words) > 1 {
			secondArgument = &words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, secondArgument)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	trimmedText := strings.Fields(loweredText)

	return trimmedText
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			callback:    commandInspect,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See all the pokemon you've caught",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
