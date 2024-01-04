package server

import (
	"htmx/src/shared/infra/env"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Serve(envConfig *env.Config) {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	port := envConfig.PORT

	e.Logger.Fatal(e.Start(":" + port))
}
