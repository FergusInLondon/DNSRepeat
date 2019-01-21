package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockLookupResolver struct {
	Calls int
}

func (mr *MockLookupResolver) LookupHost(host string) (addrs []string, err error) {
	mr.Calls++
	return []string{testTable[host], "likelywontusethis", "orthis,butwe'llsee"}, nil
}

var testTable = map[string]string{
	"google.com":   "127.0.0.1",
	"resolver.com": "192.168.0.1",
}

func TestDNSResultsAreValid(t *testing.T) {
	resolver := NewResolver(&MockLookupResolver{})

	for domain, ip := range testTable {
		addr, err := resolver.Resolve(domain)

		assert.Empty(t, err)
		assert.Equal(t, addr, ip)
	}
}

func TestResolverInterfaceIsUsed(t *testing.T) {
	lookupResolver := &MockLookupResolver{}
	resolver := NewResolver(lookupResolver)

	callNumber := 0

	for domain, _ := range testTable {
		resolver.Resolve(domain)
		callNumber++

		assert.Equal(t, callNumber, lookupResolver.Calls)
	}
}
