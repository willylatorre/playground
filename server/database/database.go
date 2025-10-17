package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the SQLite database connection and creates tables
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./adrian.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Create tables
	createTables()
}

// createTables creates the necessary database tables
func createTables() {
	coffeeTable := `
	CREATE TABLE IF NOT EXISTS coffee (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		counter INTEGER NOT NULL DEFAULT 0,
		last_update DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := DB.Exec(coffeeTable); err != nil {
		log.Fatal("Failed to create coffee table:", err)
	}

	// Insert initial row if table is empty
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM coffee").Scan(&count)
	if err != nil {
		log.Fatal("Failed to check coffee table:", err)
	}

	if count == 0 {
		_, err = DB.Exec("INSERT INTO coffee (counter) VALUES (67)")
		if err != nil {
			log.Fatal("Failed to insert initial coffee record:", err)
		}
	}

	log.Println("Database initialized successfully")
}

// Close closes the database connection
func Close() {
	if DB != nil {
		DB.Close()
	}
}
