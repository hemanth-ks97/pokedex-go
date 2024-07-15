package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/hemanth-ks97/pokedex-go/internal/pokeapi"
)

func mapb_cb(pokeclient *pokeapi.PokeClient) error {
	if cur_location_obj.Prev == "" {
		return errors.New("you have reached the first page")
	}
	//check cache and return if hit
	cached_bytes, is_present := pokeclient.Cache.Get(cur_location_obj.Prev)
	if is_present {
		cur_location_obj.Next = ""
		cur_location_obj.Prev = ""
		err := json.Unmarshal(cached_bytes, &cur_location_obj)
		if err != nil {
			return err
		}
		fmt.Println("Cache Hit!")
		cur_location_obj.PrintLocations()
		return nil
	}

	//cache miss
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

	//add to cache
	pokeclient.Cache.Add(cur_location_obj.Prev, body)
	cur_location_obj.Next = ""
	cur_location_obj.Prev = ""
	err = json.Unmarshal(body, &cur_location_obj)
	if err != nil {
		return err
	}

	// for _, locs := range cur_location_obj.Results {
	// 	fmt.Println(locs.Name)
	// }

	cur_location_obj.PrintLocations()

	return nil
}
