package main

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

// JSONDNSRequest is a wrapper for incoming DNS resolution requests.
type JSONDNSRequest struct {
	Hostname string `json:"hostname"`
}

// JSONDNSResponse is a wrapper for key/value pair responses.
type JSONDNSResponse struct {
	Hostname string `json:"hostname"`
	Address  string `json:"address"`
}

// DNSRequestHandler is simply HTTP HandlerFunc ("Handle(...)") with some
//  additional context in the form of a ResolverInterface.
type DNSRequestHandler struct {
	Resolver ResolverInterface
}

// NewDNSHandler returns a DNSRequestHandler object, and requires a valid
//  ResolverInterface
func NewDNSHandler(r ResolverInterface) *DNSRequestHandler {
	return &DNSRequestHandler{
		Resolver: r,
	}
}

// ParseHostFromRequest takes the an io.ReadCloser (i.e the Request Body) and
//  ensures it's correct, that there's a hostname property, and that the property
//  is a valid hostname for the purposes of DNS.
func ParseHostFromRequest(body io.ReadCloser) (string, error) {
	invalidPayloadError := errors.New("Invalid Payload Supplied in Request")

	if body == nil {
		return "", invalidPayloadError
	}

	var dnsRequest JSONDNSRequest
	decoder := json.NewDecoder(body)

	err := decoder.Decode(&dnsRequest)
	if err != nil || len(dnsRequest.Hostname) < 1 {
		return "", invalidPayloadError
	}

	if !govalidator.IsDNSName(dnsRequest.Hostname) {
		return "", invalidPayloadError
	}

	return dnsRequest.Hostname, nil
}

// ServeHTTP parses the initial request, ensuring the validity of the provided hostname,
//  attempts to resolve it, and returns a key/pair (as JSON) to the client.
//
// Invalid Request Payloads result in a HTTP Status Code of 400.
// Hostnames that are unable to be resolved return a HTTP Status Code of 404
// Successful resolutions return a HTTP Status Code of 200, and the key/pair.
func (drh *DNSRequestHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	host, err := ParseHostFromRequest(req.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	addr, err := drh.Resolver.Resolve(host)
	if err != nil {
		// Most likely DNS Resolution is flunky; either no records or network issues
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(&JSONDNSResponse{
		Hostname: host,
		Address:  addr,
	})

	if err != nil {
		// We need to do something, but headers have already been sent.. awkward.
	}
}
