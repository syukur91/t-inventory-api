package main

import (
	"net/http"
	"os"

	"github.com/syukur91/t-inventory-api/auth"
	"github.com/syukur91/t-inventory-api/authhandler"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func setupHandlers(e *echo.Echo, log *log.Entry) {

	r := e.Group("/:tenant/api")

	r.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello "+c.Param("tenant")+"! This is API version: "+os.Getenv("VERSION"))
	})

	authClient := auth.NewAuthCLI(log)

	authHandler := &authhandler.Handler{
		AuthClient: authClient,
		Log:        log,
	}

	authHandler.SetRoutes(r)
}
