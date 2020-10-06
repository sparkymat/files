package file

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HandlerConfig interface {
}

func Handler(cfg HandlerConfig) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusInternalServerError, "not implemented")
	}
}
