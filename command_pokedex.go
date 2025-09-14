package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	pokedex := cfg.pokedex

	if len(pokedex) == 0 {
		return errors.New("you don't have any Pokémon in your Pokedex")
	}

	fmt.Println("Your Pokedex:")
	for name := range pokedex {
		fmt.Println("-", name)
	}
	return nil
}
