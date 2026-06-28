package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *cmdConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("requires a pokemon name")
	}
	target := args[0]
	pokemon, err := cfg.client.GetPokemonStats(target)
	if err != nil {
		return fmt.Errorf("Somthing's wrong\n%v", err)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	chance := 22.0 / float64(pokemon.Base_experience)
	rand_seed := rand.Float64()
	fmt.Printf("chance to catch: %.2f %%\nrolled: %.2f\n", chance*100, rand_seed)
	if rand_seed < chance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = Pokemon{
			Name:            pokemon.Name,
			Height:          pokemon.Height,
			Weight:          pokemon.Weight,
			Base_experience: pokemon.Base_experience,
			Stats:           pokemon.GetStats(),
			Types:           pokemon.GetTypes(),
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
