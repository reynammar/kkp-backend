package handlers

import (
	"kkp-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var adminService = services.AdminService{}

func GetManifest(c *gin.Context) {
	scheduleID := c.Param("schedule_id")

	manifest, err := adminService.GetPassengerManifest(scheduleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memuat manifest penumpang"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"schedule_id":      scheduleID,
		"total_passengers": len(manifest),
		"data":             manifest,
	})
}
