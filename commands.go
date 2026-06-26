package main

import (
	"fmt"
	"os"

	"github.com/EvanHo31/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*cmdConfig) error
}

type cmdConfig struct {
	Previous string
	Next     string
	Cache    *pokecache.Cache
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
	}
	return commandMap
}

func commandExit(config *cmdConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *cmdConfig) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommand() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("-----------------------")
	return nil
}

func commandMap(config *cmdConfig) error {
	if config.Next == "" {
		return fmt.Errorf("you are on the last page")
	}
	locations, err := getLocations(config, "next")
	if err != nil {
		return err
	}
	config.Next = locations.Next
	config.Previous = locations.Previous
	printLocations(locations.Results)
	return nil
}

func commandMapB(config *cmdConfig) error {
	if config.Previous == "" {
		return fmt.Errorf("you are on the first page")
	}
	locations, err := getLocations(config, "previous")
	if err != nil {
		return err
	}
	config.Next = locations.Next
	config.Previous = locations.Previous
	printLocations(locations.Results)
	return nil
}
