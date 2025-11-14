package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "[github.com/lib/pq](https://github.com/lib/pq)"
)

// InitDB initializes and returns a database connection
func InitDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to the database")
	return db, nil
}

// CreateTables ensures the necessary tables exist
func CreateTables(db *sql.DB) error {
	userTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        name STRING NOT NULL,
        email STRING NOT NULL UNIQUE,
        password_hash STRING NOT NULL,
        role STRING NOT NULL DEFAULT 'user',
        created_at TIMESTAMPTZ DEFAULT now()
    );`

	if _, err := db.Exec(userTableSQL); err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	appointmentTableSQL := `
    CREATE TABLE IF NOT EXISTS appointments (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        user_name STRING NOT NULL,
        doctor STRING NOT NULL,
        date DATE NOT NULL,
        time STRING NOT NULL,
        reason STRING,
        status STRING NOT NULL DEFAULT 'pending',
        created_at TIMESTAMPTZ DEFAULT now()
    );`

	if _, err := db.Exec(appointmentTableSQL); err != nil {
		return fmt.Errorf("failed to create appointments table: %w", err)
	}

	log.Println("Database tables checked/created successfully")
	return nil
}
