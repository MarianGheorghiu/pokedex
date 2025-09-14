package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(locationName string) (RespLocation, error) {
	url := baseURL + "/location-area/" + locationName

	// 1. Verificăm cache-ul
	if cachedData, ok := c.cache.Get(url); ok {
		// folosim cachedData direct
		pokemonsResp := RespLocation{}
		err := json.Unmarshal(cachedData, &pokemonsResp)
		if err != nil {
			return RespLocation{}, err
		}
		return pokemonsResp, nil
	}
	// 2. Dacă nu există în cache, facem request HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}

	pokemonsResp := RespLocation{}
	err = json.Unmarshal(data, &pokemonsResp)
	if err != nil {
		return RespLocation{}, err
	}

	// 3. Adăugăm răspunsul în cache
	c.cache.Add(url, data)
	return pokemonsResp, nil
}
