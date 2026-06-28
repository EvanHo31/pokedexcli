package pokeapi

import (
	"fmt"
	"testing"
	"time"
)

func printPokemons(pokemons []pokeItem) {
	for _, pokemon := range pokemons {
		fmt.Println(pokemon.Name)
	}
}

func TestGetPokemonList(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "1",
			expected: []string{
				"tentacool",
				"tentacruel",
				"staryu",
				"magikarp",
				"gyarados",
				"wingull",
				"pelipper",
				"shellos",
				"gastrodon",
				"finneon",
				"lumineon",
			},
		},
	}
	client := Newclient(10*time.Second, 5*time.Second)
	for _, c := range cases {
		pokemons, err := client.GetPokemonList(c.input)
		if err != nil {
			t.Errorf("api error: %v", err)
			return
		}
		if len(pokemons) != len(c.expected) {
			printPokemons(pokemons)
			t.Errorf("unexpected data len: expected-%d, got-%d", len(c.expected), len(pokemons))
			return
		}
	}

}
