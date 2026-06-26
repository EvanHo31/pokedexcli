package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mux     *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: map[string]cacheEntry{},
		mux:     &sync.RWMutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c Cache) Add(key string, val []byte) {
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mux.Lock()
	c.entries[key] = newEntry
	c.mux.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	entry, ok := c.entries[key]
	c.mux.RUnlock()
	return entry.val, ok
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.entries {
			delta := time.Since(entry.createdAt)
			if delta > interval {
				delete(c.entries, key)
			}
		}
		c.mux.Unlock()
	}
}
