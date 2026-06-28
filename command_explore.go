package main

import "fmt"

func commandExplore(cfg *cmdConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("requires an area name")
	}
	area := args[0]
	pokemons, err := cfg.client.GetPokemonList(area)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", area)
	for _, pokemon := range pokemons {
		fmt.Println(pokemon.Name)
	}
	return nil
}
