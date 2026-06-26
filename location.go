package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func getLocations(config *cmdConfig, option string) (pokeLocations, error) {
	url := ""
	switch option {
	case "next":
		url = config.Next
	case "previous":
		url = config.Previous
	}
	// check cache
	cache, ok := config.Cache.Get(url)
	if ok {
		return decodeLocation(cache)
	}
	// make get request
	res, err := http.Get(url)
	if err != nil {
		return pokeLocations{}, err
	}
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return pokeLocations{}, fmt.Errorf("non-ok response - %s", res.Status)
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return pokeLocations{}, err
	}
	// store cache
	config.Cache.Add(url, buf)
	return decodeLocation(buf)
}

func decodeLocation(buf []byte) (pokeLocations, error) {
	locations := pokeLocations{}
	err := json.Unmarshal(buf, &locations)
	if err != nil {
		return pokeLocations{}, err
	}
	return locations, nil
}

func printLocations(r []pokeLocationResults) {
	for _, location := range r {
		fmt.Println(location.Name)
	}
}
