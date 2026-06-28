package main

import "fmt"

func commandMap(cfg *cmdConfig) error {
	locations, err := cfg.client.GetLocations(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = locations.Next
	cfg.Previous = locations.Previous
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(cfg *cmdConfig) error {
	if cfg.Previous == "" {
		return fmt.Errorf("you are on the first page")
	}
	locations, err := cfg.client.GetLocations(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = locations.Next
	cfg.Previous = locations.Previous
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
