package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router using Gorilla Mux
	r := mux.NewRouter()

	// Middleware to handle CORS
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

			// Handle OPTIONS method for CORS pre-flight
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Call the next handler, which can be another middleware or the final handler
			next.ServeHTTP(w, r)
		})
	}

	// Apply the CORS middleware
	r.Use(corsMiddleware)

	// Define the handler function
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Your regular handler logic
		fmt.Fprintf(w, "Hello from Go!")
	})

	// Start the server with the router
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}
