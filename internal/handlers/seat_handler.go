package handlers

import (
	"kkp-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var seatService = services.SeatService{}

func GetSeats(c *gin.Context) {
	busID := c.Param("bus_id") 
	seats, err := seatService.GetSeatsByBus(busID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data kursi"})
		return
	}
	c.JSON(http.StatusOK, seats)
}
