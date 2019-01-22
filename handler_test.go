package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEmptyPayloadToRequestHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	http.HandlerFunc(DNSRequestHandler).ServeHTTP(recorder, req) //sefgault

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestMalformedPayloadToRequestHandler(t *testing.T){
	req, _ := http.NewRequest("GET", "/", strings.NewReader("{ dewfeffef }"))
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(DNSRequestHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
