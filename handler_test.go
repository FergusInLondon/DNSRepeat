package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEmptyPayloadToRequestHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	handler, _ := create_handler()
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestMalformedJSONToRequestHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", strings.NewReader("{ dewfeffef }"))
	recorder := httptest.NewRecorder()

	handler, _ := create_handler()
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestIncorrectJSONToRequestHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", strings.NewReader("{ \"Hello\" : \"World\" }"))
	recorder := httptest.NewRecorder()

	handler, _ := create_handler()
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandlerRejectsInvalidHostnames(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", strings.NewReader("{ \"hostname\": \"willerrord/sdfsff#]]#]gfm\" }"))
	recorder := httptest.NewRecorder()

	handler, _ := create_handler()
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandlerReturnsUnknownForResolutionFailures(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", strings.NewReader("{ \"hostname\": \"willerror.com\" }"))
	recorder := httptest.NewRecorder()

	handler, resolver := create_handler()
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, 1, resolver.Calls)
	assert.Equal(t, "willerror.com", resolver.CalledWith)
	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestHandlerAcceptsCorrectDomainNames(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", strings.NewReader("{ \"hostname\": \"google.com\" }"))
	recorder := httptest.NewRecorder()

	handler, resolver := create_handler()
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, 1, resolver.Calls)
	assert.Equal(t, "google.com", resolver.CalledWith)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestHandlerRespondsWithCorrectAddress(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", strings.NewReader("{ \"hostname\": \"google.com\" }"))
	recorder := httptest.NewRecorder()

	handler, _ := create_handler()
	handler.ServeHTTP(recorder, req)

	var dnsResponse JSONDNSResponse
	decoder := json.NewDecoder(recorder.Body)
	decoder.Decode(&dnsResponse)

	assert.Equal(t, "google.com", dnsResponse.Hostname)
	assert.Equal(t, "127.0.0.1", dnsResponse.Address)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
