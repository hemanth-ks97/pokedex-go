package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/hemanth-ks97/pokedex-go/internal/pokeapi"
)

var cur_location_obj = LocationAreaObj{
	Next: "https://pokeapi.co/api/v2/location-area",
	Prev: "",
	Results: []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}{},
}

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

func exit_cb(pokeclient *pokeapi.PokeClient) error {
	os.Exit(0)
	return nil
}

func map_cb(pokeclient *pokeapi.PokeClient) error {
	res, err := pokeclient.HTTPClient.Get(cur_location_obj.Next)
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return errors.New("HTTP request failed")
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	// fmt.Printf("%s\n\n\n", body)
	cur_location_obj.Next = ""
	cur_location_obj.Prev = ""
	err = json.Unmarshal(body, &cur_location_obj)
	if err != nil {
		return err
	}

	// fmt.Printf("%+v\n", cur_location_obj)

	for _, locs := range cur_location_obj.Results {
		fmt.Println(locs.Name)
	}

	return nil
}

func mapb_cb(pokeclient *pokeapi.PokeClient) error {
	if cur_location_obj.Prev == "" {
		return errors.New("you have reached the first page")
	}

	res, err := pokeclient.HTTPClient.Get(cur_location_obj.Prev)
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return errors.New("HTTP request failed")
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	// fmt.Printf("%s\n\n\n", body)
	cur_location_obj.Next = ""
	cur_location_obj.Prev = ""
	err = json.Unmarshal(body, &cur_location_obj)
	if err != nil {
		return err
	}

	// fmt.Printf("%+v\n", cur_location_obj)

	for _, locs := range cur_location_obj.Results {
		fmt.Println(locs.Name)
	}

	return nil
}
