package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ServerRun(e *echo.Echo){
  	api := e.Group("/api")

	// Basic APi endpoint
	api.GET("/message", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, from the golang World!"})
	})

}
