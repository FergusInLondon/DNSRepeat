package main

import (
	"time"
)

type DNSCache struct {
	cache CacheInterface
}

/* @see https://github.com/patrickmn/go-cache */
type CacheInterface interface {
	Add(k string, x interface{}, d time.Duration) error
	Get(k string) (interface{}, bool)
}

func NewCache(cacheInterface CacheInterface) *DNSCache {
	return &DNSCache{
		cache: cacheInterface,
	}
}

func (c *DNSCache) Get(host string) string {
	if record, found := c.cache.Get(host); found {
		return record.(string)
	}

	return ""
}

func (c *DNSCache) Add(host, addr string) error {
	return c.cache.Add(host, addr, 10*time.Minute)
}
