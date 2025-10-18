package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the SQLite database connection and creates tables
func InitDB(dbPath string, maxOpenConns, maxIdleConns int) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Create tables
	if err = createTables(db); err != nil {
		return nil, err
	}
	
	log.Printf("Database initialized successfully at %s", dbPath)
	return db, nil
}

// createTables creates the necessary database tables
func createTables(db *sql.DB) error {
	coffeeTable := `
	CREATE TABLE IF NOT EXISTS coffee (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		counter INTEGER NOT NULL DEFAULT 0,
		last_update DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(coffeeTable); err != nil {
		return err
	}

	// Insert initial row if table is empty
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM coffee").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		_, err = db.Exec("INSERT INTO coffee (counter) VALUES (67)")
		if err != nil {
			return err
		}
	}

	return nil
}

// Close closes the database connection
func Close() {
	if DB != nil {
		DB.Close()
	}
}
