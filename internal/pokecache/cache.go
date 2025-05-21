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
	cacheEntry map[string]cacheEntry
	mu         *sync.Mutex
	interval   time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheEntry: make(map[string]cacheEntry),
		mu:         &sync.Mutex{},
		interval:   interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newCacheEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheEntry[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cacheEntry[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		currentTime := time.Now()
		for key, entry := range c.cacheEntry {
			if currentTime.After(entry.createdAt.Add(c.interval)) {
				delete(c.cacheEntry, key)
			}
		}
		c.mu.Unlock()
	}
}
