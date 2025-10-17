package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"playground-server/models"
)

// HealthCheck returns server health status
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Server is running",
	})
}

// GetItems returns all items
func GetItems(c *gin.Context) {
	items := models.GetAllItems()
	c.JSON(http.StatusOK, gin.H{
		"data": items,
	})
}

// GetItem returns a single item by ID
func GetItem(c *gin.Context) {
	id := c.Param("id")

	item, exists := models.GetItemByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Item not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}

// CreateItem creates a new item
func CreateItem(c *gin.Context) {
	var input models.CreateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	item := models.Item{
		ID:   uuid.New().String(),
		Name: input.Name,
		Data: input.Data,
	}

	models.CreateItem(item)

	c.JSON(http.StatusCreated, gin.H{
		"data": item,
	})
}

// UpdateItem updates an existing item
func UpdateItem(c *gin.Context) {
	id := c.Param("id")

	var input models.UpdateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	item, exists := models.GetItemByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Item not found",
		})
		return
	}

	// Update fields if provided
	if input.Name != nil {
		item.Name = *input.Name
	}
	if input.Data != nil {
		item.Data = *input.Data
	}

	models.UpdateItem(item)

	c.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}

// DeleteItem deletes an item by ID
func DeleteItem(c *gin.Context) {
	id := c.Param("id")

	exists := models.DeleteItem(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Item not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item deleted successfully",
	})
}
