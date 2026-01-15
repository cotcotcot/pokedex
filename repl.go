package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"pokedex/internal/pokeapi"

)

type config struct {
	pokeapiClient    pokeapi.Client
	caught map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
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

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, words[1:])
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
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
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
		"inspect": {
			name:        "inspect",
			description: "inspect a caught pokemon",
			callback:    commandInspect,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokeman by its name ",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "explore",
			description: "Explore a given location area ",
			callback:    commandPokedex,
		},
		"explore": {
			name:        "explore",
			description: "Explore a given location area ",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
