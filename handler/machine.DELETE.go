package handler

import (
    "myapp/model"
    "net/http"
		"strconv"

    "github.com/labstack/echo/v4"
)

func MachineDELETE(c echo.Context) error {
	recipeIdStr := c.Param("recipe_id")

	recipeId, err := strconv.Atoi(recipeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid recipeId format",
		})
	}

	err = model.DeleteRecipe(recipeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to delete machine",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Machine deleted successfully",
	})
}