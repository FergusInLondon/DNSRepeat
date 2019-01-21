package main

type LookupResolver interface {
	LookupHost(host string) (addrs []string, err error)
}

type Resolver struct {
	LookupResolver LookupResolver
}

func NewResolver(resolver LookupResolver) *Resolver {
	return &Resolver{
		LookupResolver: resolver,
	}
}

func (r *Resolver) Resolve(domain string) (string, error) {
	addrs, err := r.LookupResolver.LookupHost(domain)

	if err != nil {
		return "", err
	}

	return addrs[0], nil
}
