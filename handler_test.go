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

	handler, _ := create_handler()
	handler.Handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestMalformedJSONToRequestHandler(t *testing.T){
	req, _ := http.NewRequest("GET", "/", strings.NewReader("{ dewfeffef }"))
	recorder := httptest.NewRecorder()

	handler, _ := create_handler()
	handler.Handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestIncorrectJSONToRequestHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", strings.NewReader("{ \"Hello\" : \"World\" }"))
	recorder := httptest.NewRecorder()

	handler, _ := create_handler()
	handler.Handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandlerRejectsInvalidHostnames(t *testing.T) {

}

func TestHandlerReturnsUnknownForResolutionFailures(t *testing.T) {

}

func TestHandlerAcceptsCorrectDomainNames(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", strings.NewReader("{ \"hostname\": \"google.com\" }"))
	recorder := httptest.NewRecorder()

	handler, resolver := create_handler()
	handler.Handler.ServeHTTP(recorder, req)

	assert.Equal(t, 1, resolver.Calls)
	assert.Equal(t, "google.com", resolver.CalledWith)
	assert.Equal(t, http.StatusOK, recorder.Code)
}