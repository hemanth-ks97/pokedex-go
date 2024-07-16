package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/hemanth-ks97/pokedex-go/internal/pokeapi"
)

const explore_baseURL = "https://pokeapi.co/api/v2/location-area/"

func explore_cb(pokeclient *pokeapi.PokeClient, args []string) error {
	//input checks
	if len(args) < 1 {
		return errors.New("invalid usage - need additional arguments")
	}
	if len(args) > 1 {
		return errors.New("invalid usage - received more arguments than expected")
	}
	//cache hit
	cached_res, is_present := pokeclient.Cache.Get(explore_baseURL + args[0])
	if is_present {
		fmt.Println("Cache Hit!")
		response_obj := pokeapi.LocationAreaExploreResponse{}
		err := json.Unmarshal(cached_res, &response_obj)
		if err != nil {
			return err
		}

		response_obj.PrintPokemon()
		return nil
	}
	//cache miss
	fmt.Println("Cache Miss!")
	//http req
	res, err := pokeclient.HTTPClient.Get(explore_baseURL + args[0])
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
	//add to cache
	pokeclient.Cache.Add(explore_baseURL+args[0], body)
	//json unmarshalling
	response_obj := pokeapi.LocationAreaExploreResponse{}
	err = json.Unmarshal(body, &response_obj)
	if err != nil {
		return err
	}

	response_obj.PrintPokemon()

	return nil
}
