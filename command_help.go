package main

import "fmt"

func commandHelp(cfg *cmdConfig, args []string) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommand() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("-----------------------")
	return nil
}
