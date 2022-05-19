package config

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type ViewType string

const (
	ViewList ViewType = "list"
	ViewGrid ViewType = "grid"

	SecondsInDay = 86400
	DaysInYear   = 365
)

type Session struct {
	ViewType ViewType
}

func FromSession(c echo.Context) Session {
	sess, _ := session.Get("_session", c)

	sessionConfig := Session{
		ViewType: ViewList,
	}

	if value, exists := sess.Values["viewType"]; exists {
		if stringValue, isString := value.(string); isString {
			sessionConfig.ViewType = ViewType(stringValue)
		}
	}

	return sessionConfig
}

func (s *Session) Save(c echo.Context) {
	sess, _ := session.Get("_session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   DaysInYear * SecondsInDay,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}

	// ViewType
	sess.Values["viewType"] = string(s.ViewType)

	_ = sess.Save(c.Request(), c.Response())
}
