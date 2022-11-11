package repo

import "time"

type Config struct {
	Source   string        `envconfig:"DB_SOURCE" required:"true"`
	MaxIdle  int           `envconfig:"DB_MAX_IDLE" required:"true"`
	IdleTime time.Duration `envconfig:"DB_MAX_IDLE_TIME" required:"true"`
	Lifetime time.Duration `envconfig:"DB_LIFETIME" required:"true"`
	PoolSize int           `envconfig:"DB_POOLSIZE" required:"true"`
}
