package handler

import (
	"myapp/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func MachineRecipeNameGET(c echo.Context) error {
	roomIdStr := c.Param("room_id")
    
    // idをstringからintに変換
	roomId, err := strconv.Atoi(roomIdStr)
	if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error": "Invalid recipe ID format",
			})
	}

	machine, err := model.FindMachineRecipe(roomId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "error": "Internal Server Error",
        })
	}

	return c.JSON(http.StatusCreated, machine)
}