package pokeapi

import (
	"fmt"
	"testing"
	"time"
)

func printDetails(pokedex pokeDetails) {
	fmt.Printf("name: %s\nheight: %d\nweight: %d\nbase_exp: %d\n", pokedex.Name, pokedex.Height, pokedex.Weight, pokedex.Base_experience)
}

func TestGetPokeDetails(t *testing.T) {
	cases := []struct {
		input    string
		expected pokeDetails
	}{
		{
			input:    "1",
			expected: pokeDetails{},
		},
	}
	client := Newclient(10*time.Second, 5*time.Second)
	for _, c := range cases {
		details, err := client.GetPokemonStats(c.input)
		if err != nil {
			t.Errorf("failed api: %v", err)
		}
		if c.expected.Name != details.Name {
			t.Errorf("expect: %v\nactual: %v", c.expected.Name, details.Name)
			printDetails(details)
		}
	}
}
