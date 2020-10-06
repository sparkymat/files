package file

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/files/view"
)

type ListHandlerConfig interface {
	RootFolder() string
}

func ListHandler(cfg ListHandlerConfig) func(echo.Context) error {
	return func(c echo.Context) error {
		folderPath := filepath.Join(cfg.RootFolder(), c.Request().URL.Path)

		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			return c.String(http.StatusNotFound, "no such file or directory")
		}

		html := view.Layout(fmt.Sprintf("files"), view.List(folderPath))
		return c.HTML(http.StatusOK, html)
	}
}
