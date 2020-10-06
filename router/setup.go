package router

import (
	"crypto/subtle"

	"github.com/kr/pretty"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/files/internal/file"
)

type Config interface {
	Username() string
	Password() string
}

func Setup(router *echo.Echo, config Config) {
	router.Use(middleware.BasicAuth(func(username string, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte(config.Username())) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(config.Password())) == 1 {
			return true, nil
		}

		return false, nil
	}))
	router.Use(middleware.Logger())

	router.Static("/js", "public/js")
	router.Static("/css", "public/css")

	router.GET("*", file.Handler(config))

	pretty.Log(router.Routes())
}
