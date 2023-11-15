package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestRootEndpoint is a table-driven test for the root endpoint of your HTTP server
func TestRootEndpoint(t *testing.T) {
	testCases := []struct {
		name           string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		{"GET Request", "GET", http.StatusOK, "Hello from Go!"},
		{"OPTIONS Request", "OPTIONS", http.StatusOK, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			// Simulate the middleware and handler behavior
			corsMiddleware := func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					// CORS headers setup
					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
					w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

					// Handling OPTIONS method
					if r.Method == "OPTIONS" {
						w.WriteHeader(http.StatusOK)
						return
					}
					next.ServeHTTP(w, r)
				})
			}

			// Define the handler function
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello from Go!")
			})

			// Apply the middleware to the handler
			finalHandler := corsMiddleware(handler)
			finalHandler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("expected status %v; got %v", tc.expectedStatus, status)
			}

			// Check the response body
			if rr.Body.String() != tc.expectedBody {
				t.Errorf("expected body %v; got %v", tc.expectedBody, rr.Body.String())
			}
		})
	}
}
