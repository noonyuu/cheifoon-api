package handler

import (
	"myapp/model"
	"net/http"

	"os"
	"fmt"
	"io"

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

	model.CreateRecipe(&model.Recipe{UserId: 1, RecipeName: "オムライス", MenuImage: "11.jpg"})
	reset("omuraisu","11")
	model.CreateRecipe(&model.Recipe{UserId: 1, RecipeName: "ハンバーグ", MenuImage: "12.jpg"})
	reset("hamburg","12")

	model.CreateMenu(&model.Menu{UserId: 1, RecipeId: 1,SeasoningId: 1,TableSpoon: 1,TeaSpoon: 1})
	model.CreateMenu(&model.Menu{UserId: 1, RecipeId: 1,SeasoningId: 2,TableSpoon: 1,TeaSpoon: 2})
	model.CreateMenu(&model.Menu{UserId: 1, RecipeId: 2,SeasoningId: 3,TableSpoon: 2,TeaSpoon: 1})

	return c.JSON(http.StatusOK, "reset")
}

func reset(recipeName string, fileName string) {
	saveDir := "/public/uploads/"
	// 画像ファイルのパス
	imagePath := "./images/" + recipeName + ".jpg"

	// 画像ファイルを開いて読み込む
	imageFile, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("画像ファイルの読み込みに失敗しました:", err)
		return
	}
	defer imageFile.Close()

	// 保存ディレクトリを作成
	err = os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		fmt.Println("保存ディレクトリの作成に失敗しました:", err)
		return
	}

	// 画像ファイルを保存
	outFile, err := os.Create(saveDir + fileName + ".jpg")
	if err != nil {
		fmt.Println("保存ファイルの作成に失敗しました:", err)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, imageFile)
	if err != nil {
		fmt.Println("画像ファイルのコピーに失敗しました:", err)
		return
	}
}
