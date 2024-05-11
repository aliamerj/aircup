package main

import (
	"fmt"

	"github.com/aliamerj/aircup/client"
	"github.com/aliamerj/aircup/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new echo server
	e := echo.New()

	// Add standard middleware
	e.Use(middleware.Logger())

	// Setup the client (frontend) handlers to service vite static assets
	client.ClientRun(e)
	// Setup the server (backend) handlers
	server.ServerRun(e)
	// Setup the API Group
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 3000)))
}
