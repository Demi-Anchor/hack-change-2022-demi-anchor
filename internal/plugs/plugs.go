package plugs

import "demi-anchor/internal/models"

type plugs struct {
}

func New() (*plugs, error) {
	return &plugs{}, nil
}

func (p *plugs) CreateUser(u models.User) error {
	return nil
}
