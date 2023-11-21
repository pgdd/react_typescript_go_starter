package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	// Fetch database connection details from environment variables
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT") // Optional, set a default if not provided

	// Construct DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Set up connection pool here if needed

	return db, nil
}
