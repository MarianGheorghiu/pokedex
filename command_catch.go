package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	// Obține detaliile Pokémonului
	pokemonDetails, err := cfg.pokeapiClient.PokemonDetails(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonDetails.Name)

	// Calculăm șansa de captură invers proporțională cu BaseExperience
	catchChance := 100 - pokemonDetails.BaseExperience
	if catchChance < 5 {
		catchChance = 5 // minim 5% șansă
	}
	if catchChance > 95 {
		catchChance = 95 // maxim 95% șansă
	}

	// Generează un număr random între 0 și 99 folosind generatorul global (fără Seed)
	roll := rand.Intn(100)

	if roll < catchChance {
		// Pokémon prins
		fmt.Printf("You caught %v!\n", pokemonDetails.Name)
		cfg.pokedex[pokemonDetails.Name] = pokemonDetails
	} else {
		// Pokémon scăpat
		fmt.Printf("%v escaped!\n", pokemonDetails.Name)
	}

	return nil
}
