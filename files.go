package main

//go:generate go run github.com/valyala/quicktemplate/qtc -dir=view

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/files/config"
	"github.com/sparkymat/files/router"
)

func main() {
	appConfig, err := config.New()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	router.Setup(e, appConfig)

	if err = e.Start(fmt.Sprintf(":%d", appConfig.Port())); err != nil {
		panic(err)
	}
}
