package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hemanth-ks97/pokedex-go/internal/pokeapi"
)

func startRepl(client *pokeapi.PokeClient) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		command, exists := GetCommands()[input]
		if exists {
			err := command.callback(client)

			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Command %q doesn't exist\n", input)
		}
	}
}
