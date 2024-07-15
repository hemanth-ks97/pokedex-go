package main

import (
	"fmt"

	"github.com/hemanth-ks97/pokedex-go/internal/pokeapi"
)

func help_cb(pokeclient *pokeapi.PokeClient) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
