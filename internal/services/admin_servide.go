package services

import (
	"kkp-backend/internal/config"
	"kkp-backend/internal/models"
)

type AdminService struct{}

func (s *AdminService) GetPassengerManifest(scheduleID string) ([]models.Booking, error) {
	var manifest []models.Booking

	err := config.DB.
		Where("schedule_id = ? AND booking_status = ?", scheduleID, "Paid").
		Preload("User"). // Menarik detail entitas penumpang (Nama, Email)
		// Preload("Seat"). // Aktifkan jika relasi Seat ditambahkan
		Order("booking_date ASC").
		Find(&manifest).Error

	return manifest, err
}
