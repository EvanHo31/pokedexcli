package main

import "fmt"

func commandInspect(cfg *cmdConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("requires a pokemon name")
	}
	target := args[0]
	pokemon, ok := cfg.pokedex[target]
	if !ok {
		return fmt.Errorf("<%s> don't exist in the pokedex", target)
	}
	fmt.Printf(
		`Name: %s
Height: %d
Weight: %d
Base Exp: %d
`, pokemon.Name, pokemon.Height, pokemon.Weight, pokemon.Base_experience)
	fmt.Println("Stats:")
	for stat, value := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat, value)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t)
	}
	return nil
}
