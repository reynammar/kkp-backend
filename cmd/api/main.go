package main

import (
	"log"
	"os"

	"kkp-backend/internal/config"
	"kkp-backend/internal/handlers"
	"kkp-backend/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.SeedUser()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")
	{
		// Endpoint Publik
		api.POST("/login", handlers.Login)
		api.GET("/schedules", handlers.SearchSchedules)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/seats/:bus_id", handlers.GetSeats)
			protected.POST("/payment", handlers.ConfirmPayment)
			protected.GET("/history", handlers.GetUserHistory)
		}
		admin := protected.Group("/admin")
		{
			admin.GET("/manifest/:schedule_id", handlers.GetManifest)
		}
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
