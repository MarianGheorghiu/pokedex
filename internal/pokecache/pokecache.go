package pokecache

import (
	"sync"
	"time"
)

// cacheEntry reprezintă o intrare în cache.
// - createdAt: momentul când a fost adăugată intrarea, folosit pentru a decide când expiră.
// - val: datele brute pe care le stocăm în cache (de exemplu JSON de la un API).
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache ține toate intrările de cache într-o hartă protejată de un mutex.
// - mu: asigură că accesul la map este sigur când folosim cache-ul în mai multe goroutines.
// - entries: harta efectivă unde cheia e un string (ex: URL-ul) și valoarea este cacheEntry.
type Cache struct {
	mu      *sync.Mutex
	entries map[string]cacheEntry
}

// NewCache creează un cache nou.
// Parametru:
//   - interval: cât timp să păstrăm intrările înainte să le ștergem.
//
// Creează și pornește reapLoop() într-o goroutine pentru a șterge intrările expirate automat.
func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{}, // inițializăm mutex-ul
	}

	// pornim reapLoop într-o goroutine pentru curățarea periodică a intrărilor vechi
	go c.reapLoop(interval)

	return c
}

// Add adaugă o intrare nouă în cache.
// Folosește mutex pentru a bloca map-ul în timpul scrierii.
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(), // salvăm momentul adăugării
		val:       val,
	}
}

// Get returnează valoarea stocată în cache pentru o cheie.
// Folosește mutex pentru a bloca map-ul în timpul citirii.
// Returnează un boolean care indică dacă cheia există în cache.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	return entry.val, ok
}

// reapLoop rulează într-o goroutine și curăță periodic intrările expirate.
// Folosește un ticker cu intervalul dat pentru a declanșa reap-ul.
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		// folosim timpul curent pentru a decide ce intrări sunt prea vechi
		c.reap(time.Now().UTC(), interval)
	}
}

// reap șterge intrările mai vechi decât intervalul dat.
// - now: timpul curent
// - last: durata maximă permisă pentru o intrare
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.entries {
		// dacă intrarea a fost creată înainte de acum minus interval, o ștergem
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.entries, k)
		}
	}
}
