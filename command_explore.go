package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	pokemonsResp, err := cfg.pokeapiClient.ExploreLocation(args[0])
	if err != nil {
		return err
	}

	// afișăm pokemoni
	fmt.Printf("Exploring %v...\n", pokemonsResp.Location)
	fmt.Println("Found Pokemon:")
	for _, p := range pokemonsResp.PokemonEncounters {
		fmt.Printf("- %v\n", p.Pokemon.Name)
	}
	return nil

}
