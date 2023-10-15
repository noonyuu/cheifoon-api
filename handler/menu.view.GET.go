package handler

import (
	"myapp/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func MenuViewGET(c echo.Context) error {
	uidStr := c.Param("uid")
	ridStr := c.Param("rid")

	// id を string から int に変換
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID format",
		})
	}

	rid, err := strconv.Atoi(ridStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid recipe ID format",
		})
	}

	// FindMenu 関数を修正して int を受け取るように変更
	menus, err := model.FindMenu(&uid, &rid)
	if err != nil {
		// エラーが発生した場合のハンドリング
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, menus)
}
