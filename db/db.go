package db

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() error {
	var err error
	// Using environment variables would be better in production
	connStr := "host=172.17.0.3 port=5432 user=postgres password=postgres dbname=basic_web_api sslmode=disable"

	log.Printf("Attempting to connect to PostgreSQL at %s", connStr)

	// Try connecting to the database with retries
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		DB, err = sqlx.Connect("postgres", connStr)
		if err == nil {
			log.Printf("Successfully connected to PostgreSQL database")
			break
		}
		log.Printf("Attempt %d/%d: Failed to connect to database: %v", i+1, maxRetries, err)
		if i < maxRetries-1 { // Don't sleep after the last attempt
			time.Sleep(time.Second * 5) // Increased sleep time between attempts
		}
	}

	if err != nil {
		return fmt.Errorf("error connecting to the database after %d retries: %v", maxRetries, err)
	}

	err = createTables()
	if err != nil {
		return fmt.Errorf("error creating tables: %v", err)
	}

	return nil
}

func createTables() error {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		dob DATE NOT NULL
	);

	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		content TEXT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(schema)
	if err != nil {
		log.Printf("Error creating tables: %v", err)
		return err
	}

	return nil
}
