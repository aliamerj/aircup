package server

import (
	"net/http"

	"github.com/aliamerj/aircup/server/apk/routes"
	"github.com/aliamerj/aircup/server/internal"
	"github.com/labstack/echo/v4"
)

func ServerRun(e *echo.Echo) {
	configDB := internal.ConfigDB()

	api := e.Group("/api")
	{
		routes.ConfigRoutes(api, configDB)
	}

	// Basic APi endpoint
	api.GET("/message", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, from the golang World!"})
	})

}
