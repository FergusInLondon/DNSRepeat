package main

import (
	"github.com/pkg/errors"
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
	return []string{uncachedDomains[host], "likelywontusethis", "orthis,butwe'llsee"}, nil
}

func (mci *MockCacheInterface) Get(k string) (interface{}, bool) {
	mci.GetCalls++

	if addr, ok := cachedDomains[k]; ok {
		return addr, true
	}

	return "", false
}

type MockResolver struct {
	Calls int
	CalledWith string
}

func (mr *MockResolver) Resolve(domain string) (string, error) {
	mr.Calls++
	mr.CalledWith = domain

	if (domain == "willerror.com") {
		return "", errors.New("testing error")
	}

	return "127.0.0.1", nil
}


var uncachedDomains = map[string]string{
	"google.com":   "127.0.0.1",
	"resolver.com": "192.168.0.1",
}

var cachedDomains = map[string]string{
	"cacheddomain": "192.168.0.2",
}

func create_dns_cache() (*DNSCache, *MockCacheInterface) {
	mockCache := &MockCacheInterface{0, 0}
	dnsCache := NewCache(mockCache)

	return dnsCache, mockCache
}

func create_resolver() (*Resolver, *DNSCache, *MockCacheInterface, *MockLookupResolver) {
	dnsCache, mockCache := create_dns_cache()

	lookupResolver := &MockLookupResolver{}
	resolver := NewResolver(lookupResolver, dnsCache)

	return resolver, dnsCache, mockCache, lookupResolver
}

func create_handler() (*DNSRequestHandler, *MockResolver) {
	mockResolver := &MockResolver{0, ""}
	return NewDNSHandler(mockResolver), mockResolver
}
