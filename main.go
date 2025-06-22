package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ayusudi/catify-lite/config"
	_ "github.com/ayusudi/catify-lite/docs"
	"github.com/ayusudi/catify-lite/routes"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Catify API
// @version 1.0
// @description A simple API for cat facts and comments
// @host catify-api.onrender.com
// @BasePath /
// @schemes https
func main() {
	// Connect to TiDB
	config.ConnectDB()

	// Run auto migration
	config.Migrate()

	// Initialize Echo
	e := echo.New()

	// Init your app routes (facts, comments, etc.)
	routes.InitRoutes(e)

	// Swagger UI at /docs/*
	e.GET("/docs/*", echoSwagger.WrapHandler)

	// Redirect root / to Swagger UI
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "https://catify-api.onrender.com/docs/index.html")
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 Server running at http://localhost:" + port)
	e.Logger.Fatal(e.Start(":" + port))
}
