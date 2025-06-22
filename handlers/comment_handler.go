package handlers

import (
	"net/http"

	"github.com/ayusudi/catify-lite/config"
	"github.com/ayusudi/catify-lite/models"

	"github.com/labstack/echo/v4"
)

// POST /comments
func PostComment(c echo.Context) error {
	var comment models.Comment
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid comment data"})
	}

	if result := config.DB.Create(&comment); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to post comment"})
	}

	return c.JSON(http.StatusOK, comment)
}

// GET /facts/:id/comments
func GetComments(c echo.Context) error {
	factID := c.Param("id")

	var comments []models.Comment
	if result := config.DB.Where("fact_id = ?", factID).Order("created_at desc").Find(&comments); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to load comments"})
	}

	return c.JSON(http.StatusOK, comments)
}
