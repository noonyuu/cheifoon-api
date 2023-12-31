package handler

import (
	"myapp/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RecipeViewGET(c echo.Context) error {
    idStr := c.Param("id")
    
    // id を string から int に変換
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "error": "Invalid recipe ID format",
        })
    }

    // FindRecipe 関数を修正して int を受け取るように変更
    recipe, err := model.FindRecipe(id)
    if err != nil {
        // エラーが発生した場合のハンドリング
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "error": "Internal Server Error",
        })
    }

    return c.JSON(http.StatusOK, recipe)
}