package http

import (
	"database/sql"

	"github.com/go-redis/redis"
)

type Http struct {
	DB          *sql.DB
	RedisClient *redis.Client
}

type iHttp interface {
	Launch()
}

func New(http *Http) iHttp {
	return http
}
