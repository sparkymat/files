package file

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/files/view"
)

type HandlerConfig interface {
}

func Handler(cfg HandlerConfig) func(echo.Context) error {
	return func(c echo.Context) error {
		html := view.Layout(fmt.Sprintf("files"), view.List())
		return c.HTML(http.StatusOK, html)
	}
}
