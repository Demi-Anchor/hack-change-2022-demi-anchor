package plugs

import (
	"demi-anchor/internal/handler"
	"demi-anchor/internal/models"
	"encoding/json"
	"time"
)

type plugs struct {
}

func New() (*plugs, error) {
	return &plugs{}, nil
}

func (p *plugs) CreateUser(u models.User) error {
	return nil
}

func (p *plugs) CreateDonations(d handler.Donations) ([]byte, error) {
	d.DonationsAmount = 10
	d.ToDate = time.Now()
	d.FromDate = d.ToDate.Add(-48 * time.Hour)
	d.MoneyAmount = 10_000 * 100
	donationsData, _ := json.Marshal(d)
	return donationsData, nil
}
