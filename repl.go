package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hemanth-ks97/pokedex-go/internal/pokeapi"
)

func startRepl(client *pokeapi.PokeClient) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		// split input into cmd, args []string
		cmd_list := strings.Split(input, " ")
		command, exists := GetCommands()[cmd_list[0]]
		if exists {
			args := cmd_list[1:]
			err := command.callback(client, args)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Command %q doesn't exist\n", input)
		}
	}
}
