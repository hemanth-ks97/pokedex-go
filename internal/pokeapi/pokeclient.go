package pokeapi

import (
	"net/http"
	"time"

	"github.com/hemanth-ks97/pokedex-go/internal/pokecache"
)

type PokeClient struct {
	Cache      pokecache.Cache
	HTTPClient http.Client
}

func NewPokeClient() PokeClient {
	client := new(http.Client)
	client.Timeout = time.Minute
	return PokeClient{
		Cache:      pokecache.NewCache(),
		HTTPClient: *client,
	}
}
