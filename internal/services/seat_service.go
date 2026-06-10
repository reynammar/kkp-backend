package services

import (
	"kkp-backend/internal/config"
	"kkp-backend/internal/models"
)

type SeatService struct{}

func (s *SeatService) GetSeatsByBus(busID string) ([]models.Seat, error) {
	var seats []models.Seat
	err := config.DB.Where("bus_id = ?", busID).Find(&seats).Error
	return seats, err
}