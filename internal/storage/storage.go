package storage

import (
	"database/sql"
	"demi-anchor/pkg/errtrace"
	"github.com/rs/zerolog/log"
)

type storage struct {
	*sql.DB
}

func New(cfg *Config) (*storage, error) {
	db, err := sql.Open("postgres", cfg.Source)
	if err != nil {
		return nil, errtrace.AddTrace(err)
	}

	if err = db.Ping(); err != nil {
		return nil, errtrace.AddTrace(err)
	}

	db.SetMaxOpenConns(cfg.PoolSize)
	db.SetMaxIdleConns(cfg.MaxIdle)
	db.SetConnMaxIdleTime(cfg.IdleTime)
	db.SetConnMaxLifetime(cfg.Lifetime)

	return &storage{db}, nil
}

func (s *storage) Close() {
	if err := s.DB.Close(); err != nil {
		log.Err(errtrace.AddTrace(err)).Send()
	}
}
