package productv1controller

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

type V1Product struct {
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

type iV1Product interface {
	GetProducts(c echo.Context) error
}

func New(v1Product *V1Product) iV1Product {
	return v1Product
}
