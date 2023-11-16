package controllers_test

import (
	"bytes"
	"gorestapi/api/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCreateAccount tests the creation of an account
func TestCreateAccount(t *testing.T) {
	// Setup request body and http request
	var jsonStr = []byte(`{"IBAN":"DE89 3704 0044 0532 0130 00", "BIC":"COBADEFFXXX"}`)
	req, err := http.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateAccount)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated { // Use http.StatusCreated
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Add more checks as necessary, e.g., response body
}

// TestGetAccount tests retrieving an account
func TestGetAccount(t *testing.T) {
	// Setup http request
	req, err := http.NewRequest("GET", "/accounts/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetAccount)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code and other response attributes
	// ...
}
