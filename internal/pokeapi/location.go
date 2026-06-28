package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type pokeLocationResults struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type pokeLocations struct {
	Count    int                   `json:"count"`
	Next     string                `json:"next"`
	Previous string                `json:"previous"`
	Results  []pokeLocationResults `json:"results"`
}

func (c *Client) GetLocations(url string) (pokeLocations, error) {
	if url == "" {
		url = baseURL + "/location"
	}
	// check cache
	cache, ok := c.Cache.Get(url)
	if ok {
		return decodeLocation(cache)
	}
	// make get request
	res, err := c.HttpClient.Get(url)
	if err != nil {
		return pokeLocations{}, err
	}
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return pokeLocations{}, fmt.Errorf("non-ok response - %s [%s]", res.Status, url)
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return pokeLocations{}, err
	}
	// store cache
	c.Cache.Add(url, buf)
	loc, err := decodeLocation(buf)
	if err != nil {
		return pokeLocations{}, err
	}
	return loc, nil
}

func decodeLocation(buf []byte) (pokeLocations, error) {
	locations := pokeLocations{}
	err := json.Unmarshal(buf, &locations)
	if err != nil {
		return pokeLocations{}, err
	}
	return locations, nil
}
