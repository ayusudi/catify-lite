// main.go

package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/ayusudi/catify-lite/config"
	_ "github.com/ayusudi/catify-lite/docs"
	"github.com/ayusudi/catify-lite/routes"
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

	// Swagger route
	e.GET("/swagger/doc.json", func(c echo.Context) error {
		return c.File("docs/swagger.json")
})
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	
	e.Logger.Fatal(e.Start(":8080"))
}
