package plugs

import "demi-anchor/internal/models"

type plugs struct {
}

func New() *plugs {
	return &plugs{}
}

func (p *plugs) CreateUser(u models.User) error {
	return nil
}
