package service

import (
	"demi-anchor/internal/models"
)

type Repo interface {
	CreateUser(u models.User) error
}

type service struct {
	repo Repo
}

func New(r Repo) *service {
	return &service{repo: r}
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
	return s.repo.CreateUser(u)
}
