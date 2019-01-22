package main

import (
	"net"
	"strings"
)

// LookupResolver is simply an interface around the net.LookupHost(..) method,
//  https://golang.org/pkg/net/#LookupHost - included largely for testing purposes.
type LookupResolver interface {
	LookupHost(host string) (addrs []string, err error)
}

// ResolverInterface dictates that a Resolver must have a Resolve(...) method,
//  once again, this is an interface largely for the benefits of testing.
type ResolverInterface interface {
	Resolve(domain string) (string, error)
}

// Resolver coordinates the caching and resolution of any DNS requests, as such it
//  requires two properties: a DNSCache, and a LookupResolver.
type Resolver struct {
	LookupResolver LookupResolver
	Cache          *DNSCache
}

// NativeLookupResolver is the default (i.e non-testing) implementation of the
//  LookupResolver interface. This could perhaps be best refactored as an adapter
//  in the form of http.HandlerFunc?
type NativeLookupResolver struct {}

// LookupHost simply wraps around the native net.LookupHost method.
func (nlr *NativeLookupResolver) LookupHost(host string) (addrs []string, err error) {
	return net.LookupHost(host)
}

// NewResolver creates a new Resolver object, accepting both the LookupResolver
//  and the DNSCache.
func NewResolver(resolver LookupResolver, cache *DNSCache) *Resolver {
	return &Resolver{
		LookupResolver: resolver,
		Cache:          cache,
	}
}

// Resolve accepts a DNS Hostname (in the form of a string) and attempts to
//  resolve it, whilst also ensuring that any resolutions are cached appropriately.
func (r *Resolver) Resolve(domain string) (string, error) {

	// We need to ensure case insensitivity.
	domain = strings.ToLower(domain)
	if entry := r.Cache.Get(domain); len(entry) > 1 {
		return entry, nil
	}

	// No cache entry found, look up the domain, and cache if valid.
	addr, err := r.LookupResolver.LookupHost(domain)
	if err != nil {
		return "", err
	}

	err = r.Cache.Add(domain, addr[0])
	return addr[0], err
}
