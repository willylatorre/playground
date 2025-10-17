package models

import (
	"sync"
	"time"
)

// Item represents a basic data model
type Item struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateItemInput represents input for creating an item
type CreateItemInput struct {
	Name string `json:"name" binding:"required"`
	Data string `json:"data" binding:"required"`
}

// UpdateItemInput represents input for updating an item
type UpdateItemInput struct {
	Name *string `json:"name,omitempty"`
	Data *string `json:"data,omitempty"`
}

// In-memory database (for playground purposes)
var (
	items = make(map[string]Item)
	mutex = &sync.RWMutex{}
)

// InitDB initializes the database (placeholder for future database integration)
func InitDB() {
	// In a real application, this would connect to a database
	// For now, we'll just initialize some sample data
	createSampleData()
}

// createSampleData adds some sample items for testing
func createSampleData() {
	now := time.Now()
	sampleItems := []Item{
		{
			ID:        "sample-1",
			Name:      "Sample Item 1",
			Data:      "This is sample data for item 1",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        "sample-2",
			Name:      "Sample Item 2",
			Data:      "This is sample data for item 2",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	for _, item := range sampleItems {
		items[item.ID] = item
	}
}

// GetAllItems returns all items
func GetAllItems() []Item {
	mutex.RLock()
	defer mutex.RUnlock()

	result := make([]Item, 0, len(items))
	for _, item := range items {
		result = append(result, item)
	}
	return result
}

// GetItemByID returns an item by ID
func GetItemByID(id string) (Item, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	item, exists := items[id]
	return item, exists
}

// CreateItem adds a new item
func CreateItem(item Item) {
	mutex.Lock()
	defer mutex.Unlock()

	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	items[item.ID] = item
}

// UpdateItem updates an existing item
func UpdateItem(updatedItem Item) {
	mutex.Lock()
	defer mutex.Unlock()

	if item, exists := items[updatedItem.ID]; exists {
		updatedItem.CreatedAt = item.CreatedAt
		updatedItem.UpdatedAt = time.Now()
		items[updatedItem.ID] = updatedItem
	}
}

// DeleteItem removes an item by ID
func DeleteItem(id string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := items[id]; exists {
		delete(items, id)
		return true
	}
	return false
}
