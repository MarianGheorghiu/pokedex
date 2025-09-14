package main

import (
	"time"

	"github.com/MarianGheorghiu/pokedexcli/internal/pokeapi"
)

func main() {
	// Creează un nou client pentru PokeAPI cu timeout de 5 secunde.
	// Acest client va fi folosit pentru toate request-urile HTTP către API.
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	// Construim config-ul aplicației.
	// - pokeapiClient: clientul care face request-uri la PokeAPI
	// - nextLocationsURL și prevLocationsURL sunt inițial nil (nu am încă paginare)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	// Pornim bucla principală (REPL-ul), unde utilizatorul poate introduce comenzi.
	startRepl(cfg)
}
