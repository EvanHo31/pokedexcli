package pokeapi

import "encoding/json"

type pokeStat struct {
	Base_stat int      `json:"base_stat"`
	Stat      pokeItem `json:"stat"`
}

type pokeType struct {
	Type pokeItem `json:"type"`
}

type pokeDetails struct {
	Id              int        `json:"id"`
	Name            string     `json:"name"`
	Base_experience int        `json:"base_experience"`
	Height          int        `json:"height"`
	Weight          int        `json:"weight"`
	Stats           []pokeStat `json:"stats"`
	Types           []pokeType `json:"types"`
}

func (c *Client) GetPokemonStats(name string) (pokeDetails, error) {
	url := baseURL + "/pokemon/" + name
	cache, ok := c.Cache.Get(url)
	if !ok {
		//make get request
		buf, err := c.Get(url)
		if err != nil {
			return pokeDetails{}, err
		}
		cache = buf
	}
	details := pokeDetails{}
	if err := json.Unmarshal(cache, &details); err != nil {
		return pokeDetails{}, err
	}
	return details, nil
}
