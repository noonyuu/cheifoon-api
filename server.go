package main

import (
	"myapp/handler"
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// sqlDB接続
	model.OpenDB()
	model.Migrate()
  // インスタンスを作成
  e := echo.New()

  // ミドルウェアを設定
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"http://localhost:12000"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
	// 		echo.HeaderAuthorization},
	// }))

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))

  // ルートを設定
  e.GET("/", hello) // ローカル環境の場合、http://localhost:1323/ にGETアクセスされるとhelloハンドラーを実行する
	e.GET("/reset", handler.ResetGET)
	// e.POST("/user/add", handler.UserAddPOST)
	e.POST("/seasoning/add", handler.SeasoningAddPOST)
	e.POST("/recipe/add", handler.RecipeAddPOST)
	e.GET("/recipe/view/:id", handler.RecipeViewGET)
	e.POST("/menu/add", handler.MenuAddPOST)
	e.GET("/menu/view/:uid/:rid", handler.MenuViewGET)

  // サーバーをポート番号1323で起動
  e.Logger.Fatal(e.Start(":8081"))
}

// ハンドラーを定義
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
