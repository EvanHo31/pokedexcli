package main

import "fmt"

func commandPokedex(cfg *cmdConfig, args []string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for key := range cfg.pokedex {
		fmt.Println(key)
	}
	return nil
}
