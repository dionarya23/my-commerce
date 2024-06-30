package v1routes

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

type V1Routes struct {
	Echo        *echo.Group
	DB          *sql.DB
	RedisClient *redis.Client
}

type iV1Routes interface {
	MountPing()
	MountUser()
	MountCategory()
	MountProduct()
}

func New(v1Routes *V1Routes) iV1Routes {
	return v1Routes
}
