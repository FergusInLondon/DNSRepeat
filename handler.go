package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"github.com/asaskevich/govalidator"
)

type JSONDNSRequest struct {
	Hostname string `json:"hostname"`
}

type JSONDNSResponse struct {
	Hostname string `json:"hostname"`
	Address  string `json:"address"`
}

type DNSRequestHandler struct {
	Handler http.HandlerFunc
	Resolver ResolverInterface
}

func NewDNSHandler(r ResolverInterface) *DNSRequestHandler {
	return &DNSRequestHandler{
		Resolver: r,
	}
}

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
	err = json.NewEncoder(writer).Encode(&JSONDNSResponse{
		Hostname: host,
		Address: addr,
	})

	if err != nil {
		// We need to do something, but headers have already been sent.. awkward.
	}
}
