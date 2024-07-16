package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand/v2"

	"github.com/hemanth-ks97/pokedex-go/internal/pokeapi"
)

const catch_baseURL = "https://pokeapi.co/api/v2/pokemon/"

func catch_cb(pokeclient *pokeapi.PokeClient, args []string) error {
	//input checks
	if len(args) < 1 {
		return errors.New("invalid usage - need additional arguments")
	}
	if len(args) > 1 {
		return errors.New("invalid usage - received more arguments than expected")
	}
	_, is_present := pokeclient.MyPokedex[args[0]]
	if is_present {
		return fmt.Errorf("%s has already been caught", args[0])
	}
	response_obj := pokeapi.Pokemon{}

	//cache hit
	cached_res, is_present := pokeclient.Cache.Get(catch_baseURL + args[0])
	if is_present {
		fmt.Println("Cache Hit!")
		err := json.Unmarshal(cached_res, &response_obj)
		if err != nil {
			return err
		}
	} else {
		//cache miss
		fmt.Println("Cache Miss!")
		//http req
		res, err := pokeclient.HTTPClient.Get(catch_baseURL + args[0])
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
		//json unmarshalling
		err = json.Unmarshal(body, &response_obj)
		if err != nil {
			return err
		}
		//cache adds
		pokeclient.Cache.Add(catch_baseURL+args[0], body)
	}

	fmt.Printf("Throwing a pokeball at %s ...\n", args[0])
	threshold := 30
	rand_num := rand.IntN(response_obj.BaseExperience)
	if rand_num < response_obj.BaseExperience-threshold {
		return fmt.Errorf("%s Escaped", args[0])
	}

	fmt.Printf("%s was successfully caught!\n", args[0])

	//add to user pokedex
	pokeclient.MyPokedex[args[0]] = response_obj

	return nil
}
