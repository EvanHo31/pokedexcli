package pokeapi

import (
	"fmt"
	"io"
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

func (c *Client) Get(url string) ([]byte, error) {
	// make get request
	res, err := c.HttpClient.Get(url)
	if err != nil {
		return []byte{}, err
	}
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("non-ok response - %s [%s]", res.Status, url)
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	// store cache
	c.Cache.Add(url, buf)
	return buf, nil
}
