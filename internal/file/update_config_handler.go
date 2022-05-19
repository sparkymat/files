package file

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/files/config"
)

type UpdateConfigHandlerConfig interface{}

type UpdateConfigRequest struct {
	ViewType   string `form:"viewType" json:"viewType"`
	ReturnPath string `form:"returnPath" json:"returnPath"`
}

func UpdateConfigHandler(_ UpdateConfigHandlerConfig) func(echo.Context) error {
	return func(c echo.Context) error {
		var request UpdateConfigRequest
		if err := c.Bind(&request); err != nil {
			//nolint:wrapcheck
			return c.Redirect(http.StatusSeeOther, "/")
		}

		sessionConfig := config.FromSession(c)

		switch request.ViewType {
		case string(config.ViewGrid), string(config.ViewList):
			sessionConfig.ViewType = config.ViewType(request.ViewType)
		}

		sessionConfig.Save(c)

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJavaScript)

		//nolint:wrapcheck
		return c.String(http.StatusOK, fmt.Sprintf("window.location.href='%s'", request.ReturnPath))
	}
}
