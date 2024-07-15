package main

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Displays help message",
			callback:    help_cb,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    exit_cb,
		},
		"map": {
			name:        "map",
			description: "Lists the next 20 map locations in order",
			callback:    map_cb,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 location (if exists)",
			callback:    mapb_cb,
		},
	}
}