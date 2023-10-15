package handler

import (
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SeasoningAddPOST(c echo.Context) error {

	// 調味料データを取得
	addSeasoning := new(model.UserSeasoning)
	if err := c.Bind(addSeasoning); err != nil {
		c.Logger().Error(err)
		return err
	}

	// 調味料を追加
	model.AddSeasoning(addSeasoning)

	return c.JSON(http.StatusCreated, addSeasoning)
}