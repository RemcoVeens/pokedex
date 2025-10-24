// Create a cache.reapLoop() method that is called when the cache is created (by the NewCache function).
// Each time an interval (the time.Duration passed to NewCache) passes it should remove any entries that are older than the interval.
// This makes sure that the cache doesn't grow too large over time.
// For example, if the interval is 5 seconds, and an entry was added 7 seconds ago, that entry should be removed.

package internal

import (
	"sync"
	"time"
)

type Cache struct {
	values map[string]cacheEntry
	mu     sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		values: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = cacheEntry{
		time.Now(),
		val,
	}
}
func (c *Cache) Get(key string) ([]byte, bool) {
	i, err := c.values[key]
	if !err {
		return nil, false
	}
	return i.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	// Create a new ticker that fires every 'c.reapInterval'.
	ticker := time.NewTicker(interval)

	// Ensure the ticker is stopped when the goroutine exits (e.g., if we added a stop mechanism).
	// For this example, we assume the application runs indefinitely.
	defer ticker.Stop()

	// Loop forever, waiting for the ticker to fire.
	for range ticker.C {
		c.mu.Lock()

		// Calculate the cutoff time: any entry added before this time is too old.
		cutoff := time.Now().Add(-interval)

		// A slice to collect keys of expired entries.
		keysToRemove := []string{}

		// Find all expired entries.
		for key, entry := range c.values {
			if entry.createdAt.Before(cutoff) {
				keysToRemove = append(keysToRemove, key)
			}
		}

		// Remove the expired entries.
		for _, key := range keysToRemove {
			delete(c.values, key)
		}

		// Release the lock.
		c.mu.Unlock()
	}
}
