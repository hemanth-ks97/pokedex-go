package main

import (
	"os"

	"github.com/hemanth-ks97/pokedex-go/internal/pokeapi"
)

func exit_cb(pokeclient *pokeapi.PokeClient) error {
	os.Exit(0)
	return nil
}
