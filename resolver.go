package main

import "strings"

// @see https://golang.org/pkg/net/#LookupHost
type LookupResolver interface {
	LookupHost(host string) (addrs []string, err error)
}

type Resolver struct {
	LookupResolver LookupResolver
	Cache          *DNSCache
}

func NewResolver(resolver LookupResolver, cache *DNSCache) *Resolver {
	return &Resolver{
		LookupResolver: resolver,
		Cache:          cache,
	}
}

func (r *Resolver) Resolve(domain string) (string, error) {

	domain = strings.ToLower(domain)
	if entry := r.Cache.Get(domain); len(entry) > 1 {
		return entry, nil
	}

	addr, err := r.LookupResolver.LookupHost(domain)
	if err != nil {
		return "", err
	}

	err = r.Cache.Add(domain, addr[0])
	return addr[0], err
}
