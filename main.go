// main.go

package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/ayusudi/catify-lite/config"
	_ "github.com/ayusudi/catify-lite/docs"
	"github.com/ayusudi/catify-lite/routes"
)

// @title Catify Lite API
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
	e.GET("/*", echoSwagger.WrapHandler)
	
	e.Logger.Fatal(e.Start(":8080"))
}
