package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ayusudi/catify-lite/config"
	"github.com/ayusudi/catify-lite/models"

	"github.com/labstack/echo/v4"
)

// GetRandomFact godoc
// @Summary Get a random cat fact from public API
// @Tags facts
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facts/random [get]
func GetRandomFact(c echo.Context) error {
	// GET /facts/random
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch fact"})
	}
	defer resp.Body.Close()

	var data struct {
		Fact string `json:"fact"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to parse fact"})
	}

	return c.JSON(http.StatusOK, echo.Map{"fact": data.Fact})
}

// SaveFact godoc
// @Summary Save a cat fact to the database
// @Tags facts
// @Accept json
// @Produce json
// @Param body body struct{ Fact string `json:"fact"` } true "Fact to save"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facts [post]
func SaveFact(c echo.Context) error {
	// POST /facts
	var input struct {
		Fact string `json:"fact"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid fact data"})
	}

	if input.Fact == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Fact cannot be empty"})
	}

	// Check for duplicate
	var existing models.CatFact
	if err := config.DB.Where("fact = ?", input.Fact).First(&existing).Error; err == nil {
		return c.JSON(http.StatusConflict, echo.Map{"error": "Fact already exists"})
	}

	newFact := models.CatFact{Fact: input.Fact}
	if err := config.DB.Create(&newFact).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to save fact", "details": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Fact saved successfully",
		"data":    newFact,
	})
}

// GetSavedFacts godoc
// @Summary Get all saved cat facts from the database
// @Tags facts
// @Produce json
// @Success 200 {array} models.CatFact
// @Failure 500 {object} map[string]string
// @Router /facts [get]
func GetSavedFacts(c echo.Context) error {
	// GET /facts
	var facts []models.CatFact
	if err := config.DB.Find(&facts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to load facts",
			"cause": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, facts)
}
