package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/EvanHo31/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Second)
	config := cmdConfig{
		Previous: "",
		Next:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		Cache:    &cache,
	}
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
		err := cmd.callback(&config)
		if err != nil {
			fmt.Printf("error running command %s: %v\n", cmd.name, err)
		}

	}
}
