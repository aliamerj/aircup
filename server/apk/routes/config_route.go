package routes

import (
	"github.com/aliamerj/aircup/server/apk/controllers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ConfigRoutes(route *echo.Group, db *gorm.DB) {
	route.GET("/config", func(c echo.Context) error {
		return controllers.GetConfig(c, db)
	})
}
