package pokeapi

import (
	"net/http"
	"time"
)

type PokeClient struct {
	HTTPClient http.Client
}

func NewPokeClient() PokeClient {
	client := new(http.Client)
	client.Timeout = time.Minute
	return PokeClient{
		HTTPClient: *client,
	}
}
