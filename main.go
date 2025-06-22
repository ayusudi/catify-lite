// main.go

package main

import (
	"net/http"

	"github.com/ayusudi/catify-lite/config"
	_ "github.com/ayusudi/catify-lite/docs"
	"github.com/ayusudi/catify-lite/routes"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Catify API
// @version 1.0
// @description This is an API for cat facts and comments.
// @host localhost:8080
// @BasePath /

// @schemes http
func main() {
	config.ConnectDB()
	config.Migrate()

	e := echo.New()

	routes.InitRoutes(e)

	// In main.go
	e.GET("/docs/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})

	
	e.Logger.Fatal(e.Start(":8080"))
}
