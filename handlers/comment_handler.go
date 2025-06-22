package handlers

import (
	"net/http"
	"strconv"

	"github.com/ayusudi/catify-lite/config"
	"github.com/ayusudi/catify-lite/models"

	"github.com/labstack/echo/v4"
)

// PostComment godoc
// @Summary Post a comment on a specific cat fact
// @Tags comments
// @Accept json
// @Produce json
// @Param body body models.SaveCommentRequest true "Comment to save"
// @Success 200 {object} models.Comment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /comments [post]
func PostComment(c echo.Context) error {
	var comment models.Comment
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid comment data"})
	}

	if comment.FactID == 0 || comment.Name == "" || comment.Comment == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "All fields (fact_id, name, comment) are required"})
	}

	if err := config.DB.Create(&comment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to post comment"})
	}

	return c.JSON(http.StatusOK, comment)
}

// GetComments godoc
// @Summary Get all comments for a specific cat fact
// @Tags comments
// @Produce json
// @Param id path int true "Fact ID"
// @Success 200 {array} models.Comment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facts/comments/{id} [get]
func GetComments(c echo.Context) error {
	factIDStr := c.Param("id")
	factID, err := strconv.Atoi(factIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid fact ID"})
	}

	var comments []models.Comment
	if err := config.DB.Where("fact_id = ?", factID).Find(&comments).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to load comments"})
	}

	return c.JSON(http.StatusOK, comments)
}
