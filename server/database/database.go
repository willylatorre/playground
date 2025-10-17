package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the SQLite database connection and creates tables
func InitDB() {
	var err error
	
	// Get database path from environment variable or use default
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./adrian.db"
	}
	
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// Configure connection pool
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)
	DB.SetConnMaxIdleTime(10 * time.Minute)

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
