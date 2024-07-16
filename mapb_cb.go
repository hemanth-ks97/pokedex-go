package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/hemanth-ks97/pokedex-go/internal/pokeapi"
)

func mapb_cb(pokeclient *pokeapi.PokeClient, args []string) error {
	if len(args) != 0 {
		return errors.New("invalid usage - received more arguments than expected")
	}
	if pokeapi.Cur_location_obj.Prev == "" {
		return errors.New("you have reached the first page")
	}
	//check cache and return if hit
	cached_bytes, is_present := pokeclient.Cache.Get(pokeapi.Cur_location_obj.Prev)
	if is_present {
		pokeapi.Cur_location_obj.Next = ""
		pokeapi.Cur_location_obj.Prev = ""
		err := json.Unmarshal(cached_bytes, &pokeapi.Cur_location_obj)
		if err != nil {
			return err
		}
		fmt.Println("Cache Hit!")
		pokeapi.Cur_location_obj.PrintLocations()
		return nil
	}

	//cache miss
	res, err := pokeclient.HTTPClient.Get(pokeapi.Cur_location_obj.Prev)
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
	pokeclient.Cache.Add(pokeapi.Cur_location_obj.Prev, body)
	pokeapi.Cur_location_obj.Next = ""
	pokeapi.Cur_location_obj.Prev = ""
	err = json.Unmarshal(body, &pokeapi.Cur_location_obj)
	if err != nil {
		return err
	}

	pokeapi.Cur_location_obj.PrintLocations()

	return nil
}
