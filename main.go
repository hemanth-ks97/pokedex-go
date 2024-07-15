package main

import (
	"github.com/hemanth-ks97/pokedex-go/internal/pokeapi"
)

func main() {
	client := pokeapi.NewPokeClient()
	startRepl(&client)
}
