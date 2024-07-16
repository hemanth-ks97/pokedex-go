package main

import (
	"fmt"

	"github.com/hemanth-ks97/pokedex-go/internal/pokeapi"
)

func pokedex_cb(pokeclient *pokeapi.PokeClient, args []string) error {
	if len(pokeclient.MyPokedex) == 0 {
		return fmt.Errorf("you have not caught any pokemon yet")
	}
	fmt.Println("Your Pokedex:")
	for key := range pokeclient.MyPokedex {
		fmt.Println(" - ", key)
	}
	return nil
}
