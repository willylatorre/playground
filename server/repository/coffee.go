package repository

import (
	"database/sql"
	"playground-server/models"
)

// CoffeeRepository defines the interface for coffee data operations
type CoffeeRepository interface {
	Get() (models.Coffee, error)
	Increment() (models.Coffee, error)
	Reset() (models.Coffee, error)
}

// coffeeRepository implements CoffeeRepository
type coffeeRepository struct {
	db *sql.DB
}

// NewCoffeeRepository creates a new coffee repository instance
func NewCoffeeRepository(db *sql.DB) CoffeeRepository {
	return &coffeeRepository{db: db}
}

// Get retrieves the current coffee counter from the database
func (r *coffeeRepository) Get() (models.Coffee, error) {
	var coffee models.Coffee
	row := r.db.QueryRow("SELECT id, counter, last_update FROM coffee LIMIT 1")
	err := row.Scan(&coffee.ID, &coffee.Counter, &coffee.LastUpdate)
	return coffee, err
}

// Increment increments the coffee counter atomically
func (r *coffeeRepository) Increment() (models.Coffee, error) {
	var coffee models.Coffee
	err := r.db.QueryRow(`
		UPDATE coffee 
		SET counter = counter + 1, last_update = CURRENT_TIMESTAMP 
		WHERE id = (SELECT id FROM coffee LIMIT 1)
		RETURNING id, counter, last_update
	`).Scan(&coffee.ID, &coffee.Counter, &coffee.LastUpdate)
	
	return coffee, err
}

// Reset resets the coffee counter to 0 atomically
func (r *coffeeRepository) Reset() (models.Coffee, error) {
	var coffee models.Coffee
	err := r.db.QueryRow(`
		UPDATE coffee 
		SET counter = 0, last_update = CURRENT_TIMESTAMP 
		WHERE id = (SELECT id FROM coffee LIMIT 1)
		RETURNING id, counter, last_update
	`).Scan(&coffee.ID, &coffee.Counter, &coffee.LastUpdate)
	
	return coffee, err
}
