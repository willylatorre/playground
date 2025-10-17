package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"playground-server/handlers"
	"playground-server/models"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Configure trusted proxies (development: trust localhost only)
	r.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	// Add CORS middleware for frontend communication
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Initialize database (in-memory for playground)
	models.InitDB()

	// Routes
	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
		api.GET("/items", handlers.GetItems)
		api.POST("/items", handlers.CreateItem)
		api.GET("/items/:id", handlers.GetItem)
		api.PUT("/items/:id", handlers.UpdateItem)
		api.DELETE("/items/:id", handlers.DeleteItem)
	}

	// Start server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Fatal(r.Run(port))
}
