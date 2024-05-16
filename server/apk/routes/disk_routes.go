package routes

import (
	"github.com/aliamerj/aircup/server/apk/controllers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DiskRoutes(route *echo.Group, db *gorm.DB) {
	diskRoute := route.Group("/disk")
	//  diskRoute.Use(middlewares.JWTAuth(db))
	{
		diskRoute.GET("", func(c echo.Context) error {
			return controllers.GetDirInfo(c, db)
		})

	}
}
