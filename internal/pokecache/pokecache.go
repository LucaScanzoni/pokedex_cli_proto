package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	storage 	map[string]cacheEntry
	mux 		*sync.Mutex
}

type cacheEntry struct {
	createdAt	time.Time
	val 		[]byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		storage: 	make(map[string]cacheEntry),
		mux:		&sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.storage[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.storage[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mux.Lock()
		for k, v := range c.storage {
			// k is the string key
			// v is the cacheEntry
			if v.createdAt.Before(time.Now().Add(-interval)) {
				delete(c.storage, k)
			}
		}
		c.mux.Unlock()
	}
}