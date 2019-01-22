package main

import (
	cache2 "github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"time"
)

func main() {
	// Build Deps
	cache := NewCache(cache2.New(5*time.Minute, 10*time.Minute))
	dnsResolver := NewResolver(&NativeLookupResolver{}, cache)

	// Create Resolver/Handler
	dnsHandler := NewDNSHandler(dnsResolver)
	http.HandleFunc("/", dnsHandler.ServeHTTP)

	// Listen and Serve
	log.Fatal(http.ListenAndServe(":8080", nil))
}
