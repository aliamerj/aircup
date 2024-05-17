package routes

import (
	"github.com/aliamerj/aircup/server/apk/controllers"
	"github.com/aliamerj/aircup/server/apk/middlewares"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DiskRoutes(route *echo.Group, db *gorm.DB) {
	diskRoute := route.Group("/disk", middlewares.JWTAuthorization())
	{
		diskRoute.GET("", func(c echo.Context) error {
			return controllers.GetSavedDisk(c, db)
		})

		diskRoute.GET("/load", func(c echo.Context) error {
			return controllers.GetDirInfo(c, db)
		})

		diskRoute.GET("/check", func(c echo.Context) error {
			return controllers.CheckdirPath(c, db)
		})

		diskRoute.GET("/save", func(c echo.Context) error {
			return controllers.SaveNewDirPath(c, db)
		})

	}
}
