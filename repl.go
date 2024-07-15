package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		command, exists := GetCommands()[input]
		if exists {
			err := command.callback()

			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Command %q doesn't exist\n", input)
		}
	}
}
