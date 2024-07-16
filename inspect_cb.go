package main

import (
	"errors"
	"fmt"

	"github.com/hemanth-ks97/pokedex-go/internal/pokeapi"
)

func inspect_cb(pokeclient *pokeapi.PokeClient, args []string) error {
	//input checks
	if len(args) < 1 {
		return errors.New("invalid usage - need additional arguments")
	}
	if len(args) > 1 {
		return errors.New("invalid usage - received more arguments than expected")
	}

	pokemon, caught := pokeclient.MyPokedex[args[0]]
	if !caught {
		return fmt.Errorf("%s was not caught yet", args[0])
	}

	pokemon.Inspect()
	return nil
}
