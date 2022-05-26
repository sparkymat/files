package router

import (
	"crypto/subtle"

	"github.com/gorilla/sessions"
	"github.com/kr/pretty"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/files/internal/file"
)

type Config interface {
	AuthDisabled() bool
	Username() string
	Password() string
	RootFolder() string
	SessionSecret() string
}

func Setup(router *echo.Echo, config Config) {
	if !config.AuthDisabled() {
		router.Use(middleware.BasicAuth(func(username string, password string, c echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(username), []byte(config.Username())) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte(config.Password())) == 1 {
				return true, nil
			}

			return false, nil
		}))
	}

	router.Use(middleware.Logger())
	router.Use(session.Middleware(sessions.NewCookieStore([]byte(config.SessionSecret()))))

	router.Static("/js", "public/js")
	router.Static("/css", "public/css")
	router.Static("/fonts", "public/fonts")

	router.GET("*", file.ListHandler(config))
	router.POST("/update_config", file.UpdateConfigHandler((config)))

	pretty.Log(router.Routes())
}
