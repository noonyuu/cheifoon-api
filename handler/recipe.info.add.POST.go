package handler

import (
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RecipeInfoAddPOST(c echo.Context) error {
	var addRecipeInfos []model.RecipeInfo
	if err := c.Bind(&addRecipeInfos); err != nil {
		c.Logger().Error(err)
		return err
	}

	// 複数のメニューデータをデータベースに追加
	createdRecipeInfos := make([]model.RecipeInfo, 0, len(addRecipeInfos))
	for _, addRecipeInfo := range addRecipeInfos {
		// 各メニューごとにデータベースに追加
		model.CreateRecipeInfo(&addRecipeInfo)
		createdRecipeInfos = append(createdRecipeInfos, addRecipeInfo)
	}

	return c.JSON(http.StatusCreated, createdRecipeInfos)
}
