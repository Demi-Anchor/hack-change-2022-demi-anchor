package repository

import (
	"demi-anchor/internal/models"
	"demi-anchor/pkg/errtrace"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type repository struct {
	*sqlx.DB
}

func New(cfg *Config) (*repository, error) {
	db, err := sqlx.Connect("postgres", cfg.Source)
	if err != nil {
		return nil, errtrace.AddTrace(err)
	}

	db.SetMaxOpenConns(cfg.PoolSize)
	db.SetMaxIdleConns(cfg.MaxIdle)
	db.SetConnMaxIdleTime(cfg.IdleTime)
	db.SetConnMaxLifetime(cfg.Lifetime)

	return &repository{db}, nil
}

func (r *repository) Close() {
	if err := r.DB.Close(); err != nil {
		log.Err(errtrace.AddTrace(err)).Send()
	}
}

func (r *repository) AddDonation(d *models.Donation) error {
	if _, err := r.DB.Exec(addDonationSQL, d.StreamerID, d.Author, d.Money, d.Comment, d.Time); err != nil {
		return errtrace.AddTrace(err)
	}
	return nil
}

func (r *repository) GetDailyDonations(p *models.Period) ([]models.DailyDonation, error) {
	var d []models.DailyDonation
	if err := r.DB.Select(&d, getDailyDonationsSQL, p.FirstDate, p.LastDate); err != nil {
		return nil, errtrace.AddTrace(err)
	}
	return d, nil
}
