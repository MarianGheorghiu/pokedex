package pokeapi

import (
	"net/http"
	"time"

	"github.com/MarianGheorghiu/pokedexcli/internal/pokecache"
)

// Client este wrapper peste http.Client și cache-ul
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient creează un client cu timeout și interval pentru cache
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval), // pointer la Cache
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
