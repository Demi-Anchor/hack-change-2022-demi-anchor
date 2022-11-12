package service

import (
	"demi-anchor/internal/models"
)

type Repository interface {
	CreateUser(u models.User) error
}

type service struct {
	repository Repository
}

func New(r Repository) *service {
	return &service{repository: r}
}

func (s *service) ValidateUser(u models.User) (bool, string) {
	if u.Name == "" {
		return false, "name is empty"
	}
	if u.Password == "" {
		return false, "password is empty"
	}
	return true, ""
}

func (s *service) CreateUser(u models.User) error {
	return s.repository.CreateUser(u)
}
