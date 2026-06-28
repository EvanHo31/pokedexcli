package pokeapi

import (
	"encoding/json"
)

type pokeItem struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type pokeEncounter struct {
	Pokemon pokeItem
}

type pokeLocationArea struct {
	Id                 int             `json:"id"`
	Name               string          `json:"name"`
	Location           pokeLocation    `json:"location"`
	Pokemon_encounters []pokeEncounter `json:"pokemon_encounters"`
}

func (c *Client) GetPokemonList(area_name string) ([]pokeItem, error) {
	url := baseURL + "/location-area/" + area_name
	cache, ok := c.Cache.Get(url)
	if !ok {
		// make get request
		buf, err := c.Get(url)
		if err != nil {
			return []pokeItem{}, err
		}
		cache = buf
	}
	loc_area := pokeLocationArea{}
	if err := json.Unmarshal(cache, &loc_area); err != nil {
		return []pokeItem{}, err
	}
	list := []pokeItem{}
	for _, encounter := range loc_area.Pokemon_encounters {
		list = append(list, encounter.Pokemon)
	}
	return list, nil
}
