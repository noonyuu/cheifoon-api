package handler

import (
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RoomIDuplicationCheckGET(c echo.Context) error {
	id := c.Param("id")

	room, err := model.Duplicate(&id)
    if err != nil {
        // エラー時のハンドリング
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "error": "Internal Server Error",
        })
    }

    return c.JSON(http.StatusOK, room)
}