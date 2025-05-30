package main

import (
	"log"
	"os"

	"cqrs_demo/database"
	"cqrs_demo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	database.Connect()

	// Auto migrate models
	// if err := models.AutoMigrate(database.GetDB()); err != nil {
	// 	log.Fatal("Failed to migrate database:", err)
	// }
	gin.SetMode(gin.ReleaseMode)

	// Initialize Gin router
	router := gin.New()
	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// Setup routes
	routes.SetupRoutes(router)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(router.Run(":" + port))
}
