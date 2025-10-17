package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"playground-server/models"
)

// HealthCheck returns server health status
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Server is running",
	})
}

// GetCoffee returns the current coffee counter
func GetCoffee(c *gin.Context) {
	coffee, err := models.GetCoffee()
	if err != nil {
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
func IncrementCoffee(c *gin.Context) {
	coffee, err := models.UpdateCoffeeCounter()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to increment coffee counter",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    coffee,
		"message": "Coffee counter incremented",
	})
}
