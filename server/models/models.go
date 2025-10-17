package models

import (
	"time"
	
	"playground-server/database"
)

// Coffee represents the coffee counter model
type Coffee struct {
	ID         int       `json:"id"`
	Counter    int       `json:"counter"`
	LastUpdate time.Time `json:"last_update"`
}

// GetCoffee retrieves the current coffee counter from the database
func GetCoffee() (Coffee, error) {
	var coffee Coffee
	row := database.DB.QueryRow("SELECT id, counter, last_update FROM coffee LIMIT 1")
	err := row.Scan(&coffee.ID, &coffee.Counter, &coffee.LastUpdate)
	return coffee, err
}

// UpdateCoffeeCounter increments the coffee counter and updates the last_update timestamp
func UpdateCoffeeCounter() (Coffee, error) {
	_, err := database.DB.Exec("UPDATE coffee SET counter = counter + 1, last_update = CURRENT_TIMESTAMP WHERE id = (SELECT id FROM coffee LIMIT 1)")
	if err != nil {
		return Coffee{}, err
	}

	return GetCoffee()
}

// ResetCoffeeCounter resets the coffee counter to 0 and updates the last_update timestamp
func ResetCoffeeCounter() (Coffee, error) {
	_, err := database.DB.Exec("UPDATE coffee SET counter = 0, last_update = CURRENT_TIMESTAMP WHERE id = (SELECT id FROM coffee LIMIT 1)")
	if err != nil {
		return Coffee{}, err
	}

	return GetCoffee()
}
