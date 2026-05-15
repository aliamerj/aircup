package server

import (
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	airweb "github.com/aliamerj/aircup/web"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	s.CreateRoutes(e)

	distFS, err := fs.Sub(airweb.FS, "dist")
	if err != nil {
		panic(err)
	}

	assetsFS, err := fs.Sub(distFS, "assets")
	if err != nil {
		panic(err)
	}

	e.StaticFS("/assets", assetsFS)

	e.GET("/*", func(c *echo.Context) error {
		return c.FileFS("index.html", distFS)
	})

	return e
}
