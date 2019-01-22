package main

// @see https://golang.org/pkg/net/#LookupHost
type LookupResolver interface {
	LookupHost(host string) (addrs []string, err error)
}

type Resolver struct {
	LookupResolver LookupResolver
	Cache          *DNSCache
}

// NewLookupResolver() returns a *LookupResolver (i.e &LookupResolver)
// NewCache() returns a *DNSCache  (i.e &DNSCache{})
// NewResolver() is called like this: [
//   lookup := NewLookResolver()
//   cache  := NewCache()
//   resolver := NewResolver(lookup, cache)
// ]
// Resulting in --> cannot use cache (type *DNSCache) as type DNSCache in argument to NewResolver
// Why will NewResolver(LookupResolver, DNSCache) accept a *LookupResolver as it's first param?
// Alternatively, why wont it accept a *DNSCache for it's second param then?
// Neither arguments are specified as pointers, but the first argument accepts one anyway..?
func NewResolver(resolver LookupResolver, cache *DNSCache) *Resolver {
	return &Resolver{
		LookupResolver: resolver,
		Cache: cache,
	}
}

func (r *Resolver) Resolve(domain string) (string, error) {
	addrs, err := r.LookupResolver.LookupHost(domain)

	if err != nil {
		return "", err
	}

	return addrs[0], nil
}