package main

import (
	"time"
)

// CacheInterface describes the method signatures required by any underlying
//  caching mechanism. In this case, https://github.com/patrickmn/go-cache
type CacheInterface interface {
	Add(k string, x interface{}, d time.Duration) error
	Get(k string) (interface{}, bool)
}

// DNSCache is an adapter around our CacheInterface object.
type DNSCache struct {
	cache CacheInterface
}

// NewCache creates a new DNSCache object, it requires the underlying
//  CacheInterface to use for caching requests.
func NewCache(cacheInterface CacheInterface) *DNSCache {
	return &DNSCache{
		cache: cacheInterface,
	}
}

// Get attempts to retrieve the cache item for a given host, if none is found
//  then an empty string is returned.
func (c *DNSCache) Get(host string) string {
	if record, found := c.cache.Get(host); found {
		return record.(string)
	}

	return ""
}

// Add stores a cache item for a given host.
func (c *DNSCache) Add(host, addr string) error {
	return c.cache.Add(host, addr, 10*time.Minute)
}
