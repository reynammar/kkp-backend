package main

import (
	"log"
	"os"

	"kkp-backend/internal/config"
	"kkp-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.SeedUser()

	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/login", handlers.Login)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server KKP Backend Berjalan Sempurna!",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server berjalan di port %s", port)
	router.Run(":" + port)
}
