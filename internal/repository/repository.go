package repository

import (
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

func (r *repository) AddPayment(p *Payment) error {
	if _, err := r.DB.Exec(addPaymentSQL, p.Name, p.Money); err != nil {
		return errtrace.AddTrace(err)
	}
	return nil
}

//func (r *repository) GetPaymentByPeriod(p model.Period) {
//
//}
