package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type JSONDNSRequest struct {
	Hostname string `json:"hostname"`
}

func ParseHostFromRequest(body io.ReadCloser) (string, error) {
	invalidPayloadError := errors.New("Invalid Payload Supplied in Request")

	if body == nil {
		return "", invalidPayloadError
	}

	var dnsRequest JSONDNSRequest
	decoder := json.NewDecoder(body)

	err := decoder.Decode(&decoder)
	if err != nil || len(dnsRequest.Hostname) < 1 {
		return "", invalidPayloadError
	}

	return dnsRequest.Hostname, nil
}

func DNSRequestHandler(writer http.ResponseWriter, req *http.Request) {
	_, err := ParseHostFromRequest(req.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}
