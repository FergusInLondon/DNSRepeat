package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testTable = map[string]string{
	"google.com":   "127.0.0.1",
	"resolver.com": "192.168.0.1",
}

func TestDNSResultsAreValid(t *testing.T) {
	resolver, _, _, _ := create_resolver()

	for domain, ip := range testTable {
		addr, err := resolver.Resolve(domain)

		assert.Empty(t, err)
		assert.Equal(t, addr, ip)
	}
}

func TestResolverInterfaceIsUsed(t *testing.T) {
	resolver, _, _, lookupResolver := create_resolver()

	callNumber := 0
	for domain, _ := range testTable {
		resolver.Resolve(domain)
		callNumber++

		assert.Equal(t, callNumber, lookupResolver.Calls)
	}
}

func TestResolverUsesCache(t *testing.T) {
	/* To implement */
	assert.NotEmpty(t, nil)
}

func TestResolverDoesntLookupExistingDomainNames(t *testing.T) {
	/* To implement */
	assert.NotEmpty(t, nil)
}

func TestResolverStoresNewlyEncounteredDomainNames(t *testing.T) {
	/* To implement */
	assert.NotEmpty(t, nil)
}
