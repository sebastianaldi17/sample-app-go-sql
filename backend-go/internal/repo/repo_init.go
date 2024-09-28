package repo

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Repo struct {
	db         *sqlx.DB
	redis      *redis.Client
	defaultTTL time.Duration
}

func New(db *sqlx.DB, redis *redis.Client) *Repo {
	return &Repo{
		db:         db,
		defaultTTL: time.Minute * 5,
		redis:      redis,
	}
}
