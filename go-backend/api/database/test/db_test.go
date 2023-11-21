package database_test

import (
	"gorestapi/api/database" // Adjust the import path according to your project's structure
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Ping the database to check if the connection is actually alive
	if err = db.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}
