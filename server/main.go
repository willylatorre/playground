package main

import (
	"log"
	"strings"

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

	// API routes
	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
		api.GET("/items", handlers.GetItems)
		api.POST("/items", handlers.CreateItem)
		api.GET("/items/:id", handlers.GetItem)
		api.PUT("/items/:id", handlers.UpdateItem)
		api.DELETE("/items/:id", handlers.DeleteItem)
	}

	// Catch-all handler: serve index.html for client-side routing
	r.NoRoute(func(c *gin.Context) {
		// Only serve the Vue app for non-API routes
		if !strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.File("./dist/index.html")
		}
	})

	// Start server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Fatal(r.Run(port))
}
