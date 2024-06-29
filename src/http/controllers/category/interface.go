package categoryv1controller

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

type V1Category struct {
	DB          *sql.DB
	RedisClient *redis.Client
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type iV1Category interface {
	GetCategories(c echo.Context) error
}

func New(v1 *V1Category) iV1Category {
	return v1
}
