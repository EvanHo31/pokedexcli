package pokeapi

import (
	"net/http"
	"time"

	"github.com/EvanHo31/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	HttpClient http.Client
	Cache      pokecache.Cache
}

func Newclient(timeout, cacheInterval time.Duration) Client {
	return Client{
		HttpClient: http.Client{
			Timeout: timeout,
		},
		Cache: pokecache.NewCache(cacheInterval),
	}
}
