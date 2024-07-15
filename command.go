package main

import "github.com/hemanth-ks97/pokedex-go/internal/pokeapi"

type Command struct {
	name        string
	description string
	callback    func(client *pokeapi.PokeClient) error
}
