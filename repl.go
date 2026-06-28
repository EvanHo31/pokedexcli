package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/EvanHo31/pokedexcli/internal/pokeapi"
)

type cmdConfig struct {
	Previous string
	Next     string
	client   *pokeapi.Client
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
		err := cmd.callback(cfg)
		if err != nil {
			fmt.Printf("error running command %s: %v\n", cmd.name, err)
		}

	}
}
