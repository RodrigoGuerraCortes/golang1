package main

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

	log.Println("Response Body:", w.Body.String())

	// Check the response body
	expectedResponse := `{"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Expected body %s, got %s", expectedResponse, w.Body.String())
	}

}

func TestFormdataRequestWithCustomStruct(t *testing.T) {

	// Create a GET request for the /someJSON route
	req, _ := http.NewRequest("GET", "/getb?field_a=hello&field_b=world", nil)

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Check the response body
	expectedResponse := `{"a":{"FieldA":"hello"},"b":"world"}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Expected body %s, got %s", expectedResponse, w.Body.String())
	}

}

// Bind html checkboxes
func TestFormHandler(t *testing.T) {

	// Simulate form data
	formData := url.Values{}
	formData.Add("colors[]", "red")
	formData.Add("colors[]", "green")
	formData.Add("colors[]", "blue")

	// Create a POST request with form data
	req, _ := http.NewRequest("POST", "/formHandler", strings.NewReader(formData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Assert the response body
	expectedResponse := `{"color":["red","green","blue"]}`
	if strings.TrimSpace(w.Body.String()) != expectedResponse {
		t.Errorf("Expected body %s, got %s", expectedResponse, w.Body.String())
	}
}

// Bind query string or post data
func TestStartPage(t *testing.T) {

	// Create a multipart form
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	// Add form fields
	writer.WriteField("name", "::name::")
	writer.WriteField("address", "::address::")
	writer.WriteField("birthday", "1992-03-11")
	writer.Close()

	// Create a POST request with multipart form data
	req, _ := http.NewRequest("POST", "/testing", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	t.Log(req)

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Assert the response body
	expectedResponse := `Success: Name=::name::, Address=::address::, Birthday=1992-03-11 00:00:00 +0000 UTC`
	if strings.TrimSpace(w.Body.String()) != expectedResponse {
		t.Errorf("Expected body %s, got %s", expectedResponse, w.Body.String())
	}
}

func TestBindUri(t *testing.T) {

	// Create a GET request for the /someJSON route
	req, _ := http.NewRequest("GET", "/::nombre::/ffffffff-ffff-ffff-ffff-ffffffffffff", nil)

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Check the response body
	expectedResponse := `{"name":"::nombre::","uuid":{"ID":"ffffffff-ffff-ffff-ffff-ffffffffffff","Name":"::nombre::"}}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Expected body %s, got %s", expectedResponse, w.Body.String())
	}

}

func TestLoadTemplate(t *testing.T) {
	_, err := loadTemplate()

	t.Log(loadTemplate())

	if err != nil {
		t.Fatalf("Failed to load templates: %v", err)
	}
}
