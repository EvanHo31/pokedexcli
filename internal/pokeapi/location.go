package pokeapi

import (
	"encoding/json"
)

type pokeLocation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type pokeLocations struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []pokeLocation `json:"results"`
}

func (c *Client) GetLocations(url string) (pokeLocations, error) {
	if url == "" {
		url = baseURL + "/location"
	}
	// check cache
	cache, ok := c.Cache.Get(url)
	if !ok {
		// make get request
		buf, err := c.Get(url)
		if err != nil {
			return pokeLocations{}, err
		}
		cache = buf
	}
	loc := pokeLocations{}
	if err := json.Unmarshal(cache, &loc); err != nil {
		return pokeLocations{}, err
	}
	return loc, nil
}

// func decodeLocation(buf []byte) (pokeLocations, error) {
// 	locations := pokeLocations{}
// 	err := json.Unmarshal(buf, &locations)
// 	if err != nil {
// 		return pokeLocations{}, err
// 	}
// 	return locations, nil
// }
