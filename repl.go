package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MarianGheorghiu/pokedexcli/internal/pokeapi"
)

// config reține starea aplicației și resursele globale.
// - pokeapiClient: clientul care face request-uri la PokeAPI
// - nextLocationsURL / prevLocationsURL: linkuri pentru paginare (navigare între pagini de locații)
type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
}

// startRepl lansează bucla principală (REPL = Read, Eval, Print, Loop).
// Citește input de la utilizator, caută comanda și rulează funcția aferentă.
func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		// transformă inputul în cuvinte (toate litere mici)
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		// primul cuvânt este numele comenzii
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		// căutăm comanda în map-ul getCommands()
		command, exists := getCommands()[commandName]
		if exists {
			// dacă există, rulăm funcția asociată
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			// dacă nu există, afișăm mesaj
			fmt.Println("Unknown command")
			continue
		}
	}
}

// cleanInput normalizează textul introdus:
// - transformă totul în litere mici
// - împarte în cuvinte după spații
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

// cliCommand definește o comandă în REPL.
// - name: numele efectiv al comenzii (ex: "map")
// - description: descriere pentru ajutor
// - callback: funcția care se rulează când comanda este apelată
type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

// getCommands returnează toate comenzile suportate de REPL sub forma unui map.
// cheia = numele comenzii
// valoarea = cliCommand cu descriere + funcția callback
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore pokemons in location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all pokemons",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
