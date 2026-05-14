package handlers

import (
	"net/http"
	"os"

	"github.com/aliamerj/aircup/api/database"
	"github.com/aliamerj/aircup/config"
	"github.com/aliamerj/aircup/types"
	"github.com/labstack/echo/v5"
)

type Config struct {
	DB     database.Service
	Config config.Config
}

func New(db database.Service, config config.Config) Config {
	return Config{
		DB:     db,
		Config: config,
	}
}

func withErr(c *echo.Context, err error) error {
	status := http.StatusBadRequest

	if os.IsNotExist(err) {
		status = http.StatusNotFound
	}
	return c.JSON(status, types.ErrorResponse{
		Message: err.Error(),
	})
}
