package http

import (
	"fmt"
	"net/http"
	"os"

	"dionpamungkas.com/my-commerce/src/helpers"
	v1routes "dionpamungkas.com/my-commerce/src/http/routes/v1"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

func (i *Http) Launch() {
	e := echo.New()
	e.Validator = &helpers.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = helpers.ErrorHandler
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	basePath := "/v1"
	baseUrl := e.Group(basePath)
	baseUrl.GET("", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf("API Base Code for %s", os.Getenv("ENVIRONMENT")))
	})

	v1 := v1routes.New(&v1routes.V1Routes{
		Echo:        e.Group(basePath),
		DB:          i.DB,
		RedisClient: i.RedisClient,
	})

	v1.MountPing()
	v1.MountUser()
	v1.MountCategory()
	v1.MountProduct()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}
