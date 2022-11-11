package service

import (
	"demi-anchor/internal/models"
)

type Storage interface {
	CreateUser(u models.User) error
}

type service struct {
	storage Storage
}

func New(s Storage) *service {
	return &service{storage: s}
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
	return s.storage.CreateUser(u)
}
