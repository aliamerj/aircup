package middlewares

import (
	"errors"
	"net/http"
	"os"

	"github.com/aliamerj/aircup/server/apk/controllers"
	"github.com/aliamerj/aircup/server/internal/utility"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func JWTAuth(db *gorm.DB) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		SigningMethod: "HS256",
		TokenLookup:   "cookie:accessToken",
		ErrorHandler: func(c echo.Context, err error) error {
			if errors.Is(err, jwt.ErrTokenExpired) {
				return controllers.RefreshAccessToken(c, db)
			}
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		},
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(utility.JwtAuthClaims)
		}})
}
