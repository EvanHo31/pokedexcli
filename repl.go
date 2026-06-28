package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/EvanHo31/pokedexcli/internal/pokeapi"
)

type Pokemon struct {
	Id              int            `json:"id"`
	Name            string         `json:"name"`
	Base_experience int            `json:"base_experience"`
	Height          int            `json:"height"`
	Weight          int            `json:"weight"`
	Stats           map[string]int `json:"stats"`
	Types           []string       `json:"types"`
}

type cmdConfig struct {
	Previous string
	Next     string
	client   *pokeapi.Client
	pokedex  map[string]Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*cmdConfig, []string) error
}

func getCommand() map[string]cliCommand {
	var commandMap = map[string]cliCommand{
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
			description: "display the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "display the previous names of 20 location areas in the Pokemon world",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "take a location-area name and list all the Pokemon located there",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "try to catch a pokemon",
			callback:    commandCatch,
		},
	}
	return commandMap
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	cleaned := strings.Fields(lowered)
	return cleaned
}

func startREPL(cfg *cmdConfig) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Println("scanning error:", scanner.Err())
			continue
		}
		user_input := scanner.Text()
		cleaned_input := cleanInput(user_input)
		if len(cleaned_input) == 0 {
			continue
		}
		first_word := cleaned_input[0]
		cmd, ok := getCommand()[first_word]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback(cfg, cleaned_input[1:])
		if err != nil {
			fmt.Printf("error running command %s: %v\n", cmd.name, err)
		}

	}
}
