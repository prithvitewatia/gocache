package gocache

import (
	"maps"
	"slices"
	"sync"
	"time"
)

type Cache struct {
	mu       sync.RWMutex
	items    map[string]cacheItem
	stopChan chan struct{}
}

type cacheItem struct {
	value      interface{}
	expiration int64
}

func NewCache() *Cache {
	c := &Cache{
		items:    make(map[string]cacheItem),
		stopChan: make(chan struct{}),
	}
	go c.startCleanUp(time.Minute)
	return c
}
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	exp := int64(0)
	if ttl > 0 {
		exp = time.Now().Add(ttl).UnixNano()
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = cacheItem{
		value:      value,
		expiration: exp,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	item, found := c.items[key]
	c.mu.RUnlock()
	if !found {
		return nil, false
	}
	if item.expiration > 0 && time.Now().UnixNano() > item.expiration {
		c.mu.Lock()
		delete(c.items, key)
		c.mu.Unlock()
		return nil, false
	}
	return item.value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func (c *Cache) Close() {
	close(c.stopChan)
}

func (c *Cache) startCleanUp(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.deleteExpired()
		case <-c.stopChan:
			return
		}
	}
}

func (c *Cache) deleteExpired() {
	now := time.Now().UnixNano()
	c.mu.Lock()
	for k, item := range c.items {
		if item.expiration > 0 && now > item.expiration {
			delete(c.items, k)
		}
	}
	c.mu.Unlock()
}

func (c *Cache) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return slices.Collect(maps.Keys(c.items))
}

func (c *Cache) TTL(key string) (int64, bool) {
	cacheItemValue, ok := c.items[key]
	if !ok {
		return 0, false
	}
	exp := cacheItemValue.expiration
	if exp == 0 {
		return -1, true // -1 denotes no expiry
	}
	timeRemaining := exp - time.Now().UnixNano()
	if timeRemaining <= 0 {
		c.Delete(key)
		return 0, true
	}
	return timeRemaining, true
}

func (c *Cache) FlushAll() {
	c.mu.Lock()
	defer c.mu.Unlock()
	clear(c.items)
}
