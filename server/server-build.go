package server

import (
	"log"
	"net/http"

	"github.com/aliamerj/aircup/server/apk/routes"
	"github.com/aliamerj/aircup/server/internal"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ServerRun(e *echo.Echo) {
	configDB := internal.ConfigDB()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, make sure you create .env file in the root")
		return
	}

	api := e.Group("/api")
	{
		api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:3000"}, // Allow frontend origin
			AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "X-XSRF-TOKEN"},
			AllowCredentials: true,
		}))
		//		Configure CSRF middleware
		api.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
			TokenLookup:    "cookie:XSRF-TOKEN",
			CookieName:     "XSRF-TOKEN",
			CookiePath:     "/",
			CookieSecure:   true,
			CookieHTTPOnly: false,
			CookieSameSite: http.SameSiteStrictMode,
		}))
		api.Use(middleware.SecureWithConfig(middleware.SecureConfig{
			XSSProtection:         "1; mode=block",
			ContentTypeNosniff:    "nosniff",
			XFrameOptions:         "DENY",
			HSTSMaxAge:            31536000,
			ContentSecurityPolicy: "default-src 'self'; img-src 'self' https://trusted.com; script-src 'self' https://trustedscripts.com",
		}))
		api.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
		routes.AccountRoutes(api, configDB)
	}

	// Basic APi endpoint
	api.GET("/message", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, from the golang World!"})
	})

}
