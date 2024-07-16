package pokeapi

import "fmt"

type LocationAreaObj struct {
	Count   int    `json:"count"`
	Next    string `json:"next"`
	Prev    string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

var Cur_location_obj = LocationAreaObj{
	Next: "https://pokeapi.co/api/v2/location-area",
	Prev: "",
	Results: []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}{},
}

func (l *LocationAreaObj) PrintLocations() {
	fmt.Println("Current Locations: ")
	for _, locs := range l.Results {
		fmt.Printf(" - %s\n", locs.Name)
	}
}
