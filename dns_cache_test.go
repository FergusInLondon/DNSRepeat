package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDNSCacheIsInUse(t *testing.T) {
	dnsCache, mockCache := LibTestCreateDNSCache()

	dnsCache.Get("google.com")
	assert.Equal(t, 1, mockCache.GetCalls)
}

func TestDNSCacheShouldStoreNewEntries(t *testing.T) {
	dnsCache, mockCache := LibTestCreateDNSCache()

	dnsCache.Add("google.com", "127.0.0.1")
	assert.Equal(t, 1, mockCache.AddCalls)
}
