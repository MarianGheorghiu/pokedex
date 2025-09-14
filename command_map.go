package main

import (
	"errors"
	"fmt"
)

// commandMapf (map forward) cere următoarea pagină de locații de la PokeAPI.
// Actualizează linkurile de paginare din config și afișează numele locațiilor.
func commandMapf(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	// salvăm noile linkuri pentru navigare
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	// afișăm toate locațiile primite
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// commandMapb (map backward) cere pagina anterioară de locații.
// Dacă nu există pagina anterioară, returnează o eroare.
func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
