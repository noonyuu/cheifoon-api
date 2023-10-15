package handler

import (
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RecipeAddPOST(c echo.Context) error {

	// レシピデータを取得
	addRecipe := new(model.Recipe)
	if err := c.Bind(addRecipe); err != nil {
		c.Logger().Error(err)
		return err
	}

	// レシピを追加
	err := model.CreateRecipe(addRecipe)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id": addRecipe.ID,
	})
}