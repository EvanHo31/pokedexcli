package main

import (
	"time"

	"github.com/EvanHo31/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.Newclient(10*time.Second, 5*time.Second)
	cfg := cmdConfig{
		client:  &client,
		pokedex: map[string]Pokemon{},
	}
	startREPL(&cfg)
}
