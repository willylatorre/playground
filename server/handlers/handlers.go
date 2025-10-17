package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"playground-server/repository"
)

// CoffeeHandler handles coffee-related HTTP requests
type CoffeeHandler struct {
	repo repository.CoffeeRepository
}

// NewCoffeeHandler creates a new coffee handler with dependency injection
func NewCoffeeHandler(repo repository.CoffeeRepository) *CoffeeHandler {
	return &CoffeeHandler{repo: repo}
}

// HealthCheck returns server health status
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Server is running",
	})
}

// GetCoffee returns the current coffee counter
func (h *CoffeeHandler) GetCoffee(c *gin.Context) {
	coffee, err := h.repo.Get()
	if err != nil {
		log.Printf("ERROR: Failed to get coffee counter: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get coffee counter",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": coffee,
	})
}

// IncrementCoffee increments the coffee counter
func (h *CoffeeHandler) IncrementCoffee(c *gin.Context) {
	coffee, err := h.repo.Increment()
	if err != nil {
		log.Printf("ERROR: Failed to increment coffee counter: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to increment coffee counter",
		})
		return
	}

	log.Printf("INFO: Coffee counter incremented to %d", coffee.Counter)
	c.JSON(http.StatusOK, gin.H{
		"data":    coffee,
		"message": "Coffee counter incremented",
	})
}
