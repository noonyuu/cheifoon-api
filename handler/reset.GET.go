package handler

import (
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ResetGET(c echo.Context) error {
	// tx := model.GetDB().Begin()

	model.OpenDB()
	// テーブル初期化
	model.DeleteDB()
	// // マイグレーション
	model.Migrate()

	// テストデータを作成
	model.CreateUser(&model.User{Name: "道上", Email: "miti@gmail.com", Password: "miti"})

	model.CreateSeasoning(&model.AdminSeasoning{SeasoningName: "醤油", TeaSecond: 1.2, BottleImage:"bottle.svg"})
	model.CreateSeasoning(&model.AdminSeasoning{SeasoningName: "みりん", TeaSecond: 1.2, BottleImage:"bottle.svg"})
	model.CreateSeasoning(&model.AdminSeasoning{SeasoningName: "ウスターソース", TeaSecond: 1.2, BottleImage:"bottle.svg"})
	model.CreateSeasoning(&model.AdminSeasoning{SeasoningName: "さけ", TeaSecond: 1.2, BottleImage:"bottle.svg"})
	model.CreateSeasoning(&model.AdminSeasoning{SeasoningName: "ポン酢", TeaSecond: 1.2, BottleImage:"bottle.svg"})

	model.AddSeasoning(&model.UserSeasoning{SeasoningId: 1})
	model.AddSeasoning(&model.UserSeasoning{SeasoningId: 3})
	// model.AddSeasoning(&model.UserSeasoning{SeasoningId: 1},tx)
	// model.AddSeasoning(&model.UserSeasoning{SeasoningId: 3},tx)

	model.CreateRecipe(&model.Recipe{UserId: 1, RecipeName: "オムライス", MenuImage: "omuraisu.png"})
	model.CreateRecipe(&model.Recipe{UserId: 1, RecipeName: "ハンバーグ", MenuImage: "hamburg.png"})

	model.CreateMenu(&model.Menu{UserId: 1, RecipeId: 1,SeasoningId: 1,TableSpoon: 1,TeaSpoon: 1})
	model.CreateMenu(&model.Menu{UserId: 1, RecipeId: 1,SeasoningId: 2,TableSpoon: 1,TeaSpoon: 2})
	model.CreateMenu(&model.Menu{UserId: 1, RecipeId: 2,SeasoningId: 4,TableSpoon: 2,TeaSpoon: 1})

	return c.JSON(http.StatusOK, "reset")
}
