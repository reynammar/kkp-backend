package services

import (
	"kkp-backend/internal/config"
	"kkp-backend/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentService struct{}

func (s *PaymentService) ProcessPayment(bookingID, method string) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		// 1. Update status booking menjadi 'Paid'
		if err := tx.Model(&models.Booking{}).Where("booking_id = ?", bookingID).
			Update("booking_status", "Paid").Error; err != nil {
			return err
		}

		// 2. Buat catatan pembayaran
		payment := models.Payment{
			PaymentID:     uuid.New().String(),
			BookingID:     bookingID,
			PaymentMethod: method,
			PaymentStatus: "Success",
			PaymentDate:   time.Now(),
		}
		return tx.Create(&payment).Error
	})
}
