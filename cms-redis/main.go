// main.go - CMS Redis

package main

import (
	"log"
	"os"

	"cms_redis/database"
	"cms_redis/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Khởi tạo Redis và index
	database.InitRedis()
	if err := database.EnsureLocationIndex(); err != nil {
		log.Fatalf("Không thể tạo index Redisearch: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Middleware CORS đơn giản
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

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Redis CMS server is running",
		})
	})

	// Định nghĩa routes
	routes.SetupRoutes(router)

	port := getEnv("PORT", "8081")
	log.Printf("Server starting on port %s", port)
	log.Fatal(router.Run("0.0.0.0:" + port))
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
