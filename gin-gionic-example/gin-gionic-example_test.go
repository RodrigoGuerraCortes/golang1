package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdminRouteUsingSetupRouter(t *testing.T) {
	router := setupRouter() // Call the actual router setup from main.go

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

func TestAdminRouteUsingSetupRouterBAD(t *testing.T) {
	router := setupRouter() // Call the actual router setup from main.go

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
