package services

import (
	"kkp-backend/internal/config"
	"kkp-backend/internal/models"
)

type ScheduleService struct{}

func (s *ScheduleService) GetAvailableSchedules(origin, destination, date string) ([]models.Schedule, error) {
	var schedules []models.Schedule

	err := config.DB.
		Joins("JOIN routes ON routes.route_id = schedules.route_id").
		Where("routes.origin = ? AND routes.destination = ? AND schedules.departure_date = ?",
			origin, destination, date).
		Preload("Route").
		Preload("Bus").
		Find(&schedules).Error

	return schedules, err
}
