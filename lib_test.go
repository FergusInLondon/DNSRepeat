package main

import (
	"time"
)

type MockCacheInterface struct {
	GetCalls int
	AddCalls int
}

func (mci *MockCacheInterface) Add(k string, x interface{}, d time.Duration) error {
	mci.AddCalls++
	return nil
}

type MockLookupResolver struct {
	Calls int
}

func (mr *MockLookupResolver) LookupHost(host string) (addrs []string, err error) {
	mr.Calls++
	return []string{testTable[host], "likelywontusethis", "orthis,butwe'llsee"}, nil
}

func (mci *MockCacheInterface) Get(k string) (interface{}, bool) {
	mci.GetCalls++

	if k == "isalreadycached.com" {
		return "127.0.0.1", true
	}

	return "", false
}

func create_dns_cache() (*DNSCache, *MockCacheInterface) {
	mockCache := &MockCacheInterface{0, 0}
	dnsCache := NewCache(mockCache)

	return dnsCache, mockCache
}

func create_resolver() (*Resolver, *DNSCache, *MockCacheInterface, *MockLookupResolver) {
	lookupResolver := &MockLookupResolver{}
	resolver := NewResolver(lookupResolver)

	dnsCache, mockCache := create_dns_cache()
	return resolver, dnsCache, mockCache, lookupResolver
}
