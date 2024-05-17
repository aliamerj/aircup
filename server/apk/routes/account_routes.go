package routes

import (
	"github.com/aliamerj/aircup/server/apk/controllers"
	"github.com/aliamerj/aircup/server/apk/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func AccountRoutes(route *echo.Group, db *gorm.DB) {

	authRoute := route.Group("/auth")
	{
		authRoute.Use(middleware.BodyLimit("2M"))
		authRoute.POST("/register", func(c echo.Context) error {
			return controllers.Register(c, db)
		})
		authRoute.POST("/login", func(c echo.Context) error {
			return controllers.Login(c, db)
		})

		authRoute.GET("/me", func(c echo.Context) error {
			return controllers.VerifySession(c, db)
		}, middlewares.JWTAuth(db))
		authRoute.POST("/logout", func(c echo.Context) error {
			return controllers.Logout(c, db)
		}, middlewares.JWTAuth(db))
	}

}
