package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	pokemon, ok := cfg.pokedex[args[0]]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Println(" -", stat.Stat.Name, ":", stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, ty := range pokemon.Types {
		fmt.Println(" -", ty.Type.Name)
	}
	return nil
}
