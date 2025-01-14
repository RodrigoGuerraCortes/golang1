package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router = setupRouter() // Initialize router once

func TestAdminRouteUsingSetupRouter(t *testing.T) {

	// Mock request
	payload := `{"value":"test_value"}`
	req, _ := http.NewRequest("POST", "/admin", bytes.NewBufferString(payload))
	req.SetBasicAuth("foo", "bar")
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
}

func TestAdminRouteUsingSetupRouterAuthorized(t *testing.T) {
	// Mock request
	payload := `{"value":"test_value"}`
	req, _ := http.NewRequest("POST", "/admin", bytes.NewBufferString(payload))
	req.SetBasicAuth("foo", "::error::")
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %d", w.Code)
	}
}

func TestSomeJSON(t *testing.T) {

	// Create a GET request for the /someJSON route
	req, _ := http.NewRequest("GET", "/someJSON", nil)

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Check the response body
	expectedResponse := `{"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Expected body %s, got %s", expectedResponse, w.Body.String())
	}

}
