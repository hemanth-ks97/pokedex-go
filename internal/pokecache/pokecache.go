package pokecache

import (
	"time"
)

type Cache struct {
	CacheItems map[string]cacheentry
}

func NewCache() Cache {
	c := Cache{
		CacheItems: make(map[string]cacheentry),
	}
	go c.ReapLoop()
	return c
}

func (c *Cache) Add(url string, data []byte) error {
	cache_item := new(cacheentry)
	cache_item.created_at = time.Now().UTC()
	cache_item.val = data
	c.CacheItems[url] = *cache_item
	return nil
}

func (c *Cache) Get(url string) ([]byte, bool) {
	cache_item, is_present := c.CacheItems[url]
	return cache_item.val, is_present
}

func (c *Cache) Remove(url string) error {
	delete(c.CacheItems, url)
	return nil
}

func (c *Cache) ReapLoop() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	for key, val := range c.CacheItems {
		if val.created_at.Before(time.Now().Add(-3 * time.Minute)) {
			c.Remove(key)
		}
	}
}
