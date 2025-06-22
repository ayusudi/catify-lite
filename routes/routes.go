package routes

import (
	"github.com/ayusudi/catify-lite/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	// FACT routes
	e.GET("/facts/random", handlers.GetRandomFact) // fetch from API
	e.POST("/facts", handlers.SaveFact)       // save to DB
	e.GET("/facts", handlers.GetSavedFacts)        // list saved facts

	// COMMENT routes
	e.POST("/comments", handlers.PostComment)         // comment on a fact
	e.GET("/comments/facts/:id", handlers.GetComments) // view comments for a fact
}
