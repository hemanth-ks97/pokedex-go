package main

import (
	"errors"
	"os"

	"github.com/hemanth-ks97/pokedex-go/internal/pokeapi"
)

func exit_cb(pokeclient *pokeapi.PokeClient, args []string) error {
	if len(args) != 0 {
		return errors.New("invalid usage - received more arguments than expected")
	}
	os.Exit(0)
	return nil
}
