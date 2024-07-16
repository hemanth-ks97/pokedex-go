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
		"explore": {
			name:        "explore <location-area>",
			description: "Lists the pokemon available in the <location-area>",
			callback:    explore_cb,
		},
		"catch": {
			name:        "catch <name of pokemon>",
			description: "Gives you a chance to catch the pokemon",
			callback:    catch_cb,
		},
		"inspect": {
			name:        "inspect <name of pokemon>",
			description: "Lists the stats of the caught pokemon from your Pokedex",
			callback:    inspect_cb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists the names of all caught pokemon",
			callback:    pokedex_cb,
		},
	}
}
