package main

type LocationAreaObj struct {
	Count   int    `json:"count"`
	Next    string `json:"next"`
	Prev    string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

var cur_location_obj = LocationAreaObj{
	Next: "https://pokeapi.co/api/v2/location-area",
	Prev: "",
	Results: []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}{},
}
