package handlers

import (
	"encoding/json"
	"net/http"

	"catify-lite/config"
	"catify-lite/models"

	"github.com/labstack/echo/v4"
)

// GET /facts/random
func GetRandomFact(c echo.Context) error {
	resp, err := http.Get("https://cat-fact.herokuapp.com/facts/random")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch fact"})
	}
	defer resp.Body.Close()

	var fact models.CatFact
	if err := json.NewDecoder(resp.Body).Decode(&fact); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to parse fact"})
	}

	return c.JSON(http.StatusOK, fact)
}

// POST /facts/save
func SaveFact(c echo.Context) error {
	var fact models.CatFact
	if err := c.Bind(&fact); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid fact data"})
	}

	if result := config.DB.Create(&fact); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to save fact"})
	}

	return c.JSON(http.StatusOK, fact)
}

// GET /facts
func GetSavedFacts(c echo.Context) error {
	var facts []models.CatFact
	if result := config.DB.Order("created_at desc").Find(&facts); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to load facts"})
	}

	return c.JSON(http.StatusOK, facts)
}
