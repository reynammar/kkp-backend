package services

import (
	"errors"
	"kkp-backend/internal/config"
	"kkp-backend/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookingService struct{}

func (s *BookingService) CreateBooking(userID, scheduleID, seatID string) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		// 1. Cek status kursi
		var seat models.Seat
		if err := tx.Set("gorm:query_option", "FOR UPDATE"). // Lock baris ini agar tidak di-update user lain
									Where("seat_id = ? AND seat_status = ?", seatID, "Available").
									First(&seat).Error; err != nil {
			return errors.New("kursi tidak tersedia atau sudah dipesan")
		}

		// 2. Buat data booking
		booking := models.Booking{
			BookingID:     uuid.New().String(),
			BookingDate:   time.Now(),
			BookingStatus: "Pending",
			UserID:        userID,
			ScheduleID:    scheduleID,
			SeatID:        seatID,
		}
		if err := tx.Create(&booking).Error; err != nil {
			return err
		}

		// 3. Update status kursi jadi Booked
		if err := tx.Model(&seat).Update("seat_status", "Booked").Error; err != nil {
			return err
		}

		return nil
	})
}

func (s *BookingService) GetUserHistory(userID string) ([]models.Booking, error) {
	var bookings []models.Booking

	// Preload digunakan untuk menarik data relasi agar response JSON menjadi lengkap
	err := config.DB.Where("user_id = ?", userID).
		Preload("Schedule").
		Preload("Schedule.Route").
		Preload("Schedule.Bus").
		Preload("Payment").
		Order("booking_date desc"). // Urutkan dari pesanan terbaru
		Find(&bookings).Error

	return bookings, err
}
