package pokeapi

import (
    "net/http"
    "time"

    "github.com/kevin-baik/pokedexcli/internal/pokecache"
    "github.com/kevin-baik/pokedexcli/internal/pokedex"
)

type Client struct {
    Pokedex	pokedex.Pokedex
    cache	pokecache.Cache
    httpClient	http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
    return Client{
	Pokedex:    pokedex.NewPokedex(),
	cache:	    pokecache.NewCache(cacheInterval),
	httpClient: http.Client{
	    Timeout: timeout,
	},
    }
}
