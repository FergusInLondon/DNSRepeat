package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDNSResultsAreValid(t *testing.T) {
	resolver, _, _, _ := create_resolver()

	for domain, ip := range uncachedDomains {
		addr, err := resolver.Resolve(domain)

		assert.Empty(t, err)
		assert.Equal(t, addr, ip)
	}
}

func TestDNSIgnoresCharacterCase(t *testing.T) {
	resolver, _, _, _ := create_resolver()

	for domain, ip := range uncachedDomains {
		lowerResolution, err := resolver.Resolve(domain)

		assert.Empty(t, err)
		assert.Equal(t, lowerResolution, ip)

		upperResolution, err := resolver.Resolve(strings.ToUpper(domain))
		assert.Empty(t, err)
		assert.Equal(t, upperResolution, ip)
	}
}

func TestResolverInterfaceIsUsed(t *testing.T) {
	resolver, _, _, lookupResolver := create_resolver()

	callNumber := 0
	for domain, _ := range uncachedDomains {
		resolver.Resolve(domain)
		callNumber++

		assert.Equal(t, callNumber, lookupResolver.Calls)
	}
}

func TestResolverUsesCache(t *testing.T) {
	resolver, _, mockCache, _ := create_resolver()

	cacheCalls := 0
	for domain, _ := range uncachedDomains {
		resolver.Resolve(domain)
		cacheCalls++

		assert.Equal(t, cacheCalls, mockCache.GetCalls)
	}
}

func TestResolverDoesntLookupExistingDomainNames(t *testing.T) {
	resolver, _, _, lookupResolver := create_resolver()

	callNumber := 0
	for domain, _ := range cachedDomains {
		resolver.Resolve(domain)
		callNumber++

		assert.Equal(t, 0, lookupResolver.Calls)
	}
}

func TestResolverStoresNewlyEncounteredDomainNames(t *testing.T) {
	resolver, _, mockCache, _ := create_resolver()

	cacheCalls := 0
	for domain, _ := range uncachedDomains {
		resolver.Resolve(domain)
		cacheCalls++

		assert.Equal(t, cacheCalls, mockCache.AddCalls)
	}
}
