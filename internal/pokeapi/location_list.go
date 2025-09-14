package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations face un request HTTP către PokeAPI pentru a obține o pagină de locații.
//   - pageURL: dacă e nil -> începe de la prima pagină (baseURL + "/location-area")
//     dacă e setat -> folosește acel URL (pentru paginare).
//
// Returnează un RespShallowLocations cu datele JSON decodificate.
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// 1. Verificăm cache-ul
	if cachedData, ok := c.cache.Get(url); ok {
		// folosim cachedData direct
		// json.Unmarshal(cachedData, &locationsResp)
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(cachedData, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}
	// 2. Dacă nu există în cache, facem request HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// 3. Adăugăm răspunsul în cache
	c.cache.Add(url, data)
	return locationsResp, nil
}
