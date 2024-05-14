package routes

import (
	"errors"
	"os"

	"github.com/aliamerj/aircup/server/apk/controllers"
	"github.com/aliamerj/aircup/server/internal/utility"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
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

		authRoute.Use(echojwt.WithConfig(echojwt.Config{
			SigningKey:    []byte(os.Getenv("JWT_SECRET")),
			SigningMethod: "HS256",
			TokenLookup:   "cookie:accessToken",
			ErrorHandler: func(c echo.Context, err error) error {
				if errors.Is(err, jwt.ErrTokenExpired) {
					return controllers.VerifySession(c, db)
				}
				return err
			},
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(utility.JwtAuthClaims)
			}}))
		authRoute.GET("/me", func(c echo.Context) error {
			return controllers.VerifySession(c, db)
		})
		authRoute.POST("/logout", func(c echo.Context) error {
			return controllers.Logout(c, db)
		})
	}

}
