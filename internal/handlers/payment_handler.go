package handlers

import (
	"net/http"
	"kkp-backend/internal/services"
	"github.com/gin-gonic/gin"
)

var paymentService = services.PaymentService{}

func ConfirmPayment(c *gin.Context) {
	var req struct {
		BookingID string `json:"booking_id"`
		Method    string `json:"method"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak lengkap"})
		return
	}

	if err := paymentService.ProcessPayment(req.BookingID, req.Method); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses pembayaran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pembayaran berhasil, tiket sudah aktif!"})
}