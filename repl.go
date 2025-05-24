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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "list 20 location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "list previous location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "list pokemons encounters from a specific area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch a specific pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "display pokemon information",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "display all caught pokemons",
			callback:    commandPokedex,
		},
	}
}
