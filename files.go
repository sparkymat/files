package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/files/config"
	"github.com/sparkymat/files/router"
)

func main() {
	appConfig := config.New()
	e := echo.New()

	router.Setup(e, appConfig)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.Port())))
}
