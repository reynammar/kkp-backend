package handlers

import (
	"kkp-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var scheduleService = services.ScheduleService{}

func SearchSchedules(c *gin.Context) {
	origin := c.Query("origin")
	destination := c.Query("destination")
	date := c.Query("date")

	schedules, err := scheduleService.GetAvailableSchedules(origin, destination, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil jadwal"})
		return
	}

	c.JSON(http.StatusOK, schedules)
}
