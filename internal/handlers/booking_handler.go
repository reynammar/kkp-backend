package handlers

import (
	"kkp-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var bookingService = services.BookingService{}

func CreateBooking(c *gin.Context) {
	// Ambil userID dari token JWT (asumsi sudah diset di context oleh middleware)
	userID, _ := c.Get("user_id")

	var req struct {
		ScheduleID string `json:"schedule_id"`
		SeatID     string `json:"seat_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	err := bookingService.CreateBooking(userID.(string), req.ScheduleID, req.SeatID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking berhasil dibuat!"})
}

func GetUserHistory(c *gin.Context) {
	// Ambil userID dari token JWT yang sudah di-set oleh middleware
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Sesi pengguna tidak valid"})
		return
	}

	bookings, err := bookingService.GetUserHistory(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil riwayat pesanan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil riwayat pesanan",
		"data":    bookings,
	})
}
