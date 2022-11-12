package service

import (
	"demi-anchor/internal/handler"
)

type Repository interface {
	CreateUser(u handler.User) error
	CreateDonations(d handler.Donations) ([]byte, error)
}

type service struct {
	repository Repository
}

func (s *service) CreateDonations(d handler.Donations) ([]byte, error) {
	return s.repository.CreateDonations(d)
}

func (s *service) ValidateDonation(d handler.Donations) (bool, string) {
	if d.Name == "" {
		return false, "name is empty"
	}
	if d.FromDate.After(d.ToDate) {
		return false, "from date is older than to"
	}
	if d.MoneyAmount < 0 || d.DonationsAmount < 0 {
		return false, "amount is less than zero"
	}

	return true, ""
}

func New(r Repository) *service {
	return &service{repository: r}
}

func (s *service) ValidateUser(u handler.User) (bool, string) {
	if u.Name == "" {
		return false, "name is empty"
	}
	if u.Password == "" {
		return false, "password is empty"
	}
	return true, ""
}

func (s *service) CreateUser(u handler.User) error {
	return s.repository.CreateUser(u)
}
