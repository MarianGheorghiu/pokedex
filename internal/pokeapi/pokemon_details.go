package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonDetails(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	if cachedData, ok := c.cache.Get(url); ok {
		// folosim cachedData direct
		pokemon := Pokemon{}
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	// 2. Dacă nu există în cache, facem request HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	// 3. Adăugăm răspunsul în cache
	c.cache.Add(url, data)
	return pokemon, nil
}
