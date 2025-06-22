package main

import (
	"github.com/ayusudi/catify-lite/config"
	"github.com/ayusudi/catify-lite/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.ConnectDB()
	config.Migrate()
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
