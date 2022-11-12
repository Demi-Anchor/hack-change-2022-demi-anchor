package service

import (
	"demi-anchor/internal/models"
)

type Repository interface {
	AddDonation(d *models.Donation) error
	GetDailyDonations(p *models.Period) ([]models.DailyDonation, error)
}

type service struct {
	repository Repository
}

func New(r Repository) *service {
	return &service{repository: r}
}

func (s *service) AddDonation(d *models.Donation) error {
	return s.repository.AddDonation(d)
}

func (s *service) GetDailyDonations(p *models.Period) ([]models.DailyDonation, error) {
	return s.repository.GetDailyDonations(p)
}
