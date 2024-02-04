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

		// CORSの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Reactアプリケーションのアドレス
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

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
  e.GET("/", hello)
	e.GET("/reset", handler.ResetGET)
	// e.POST("/user/add", handler.UserAddPOST)
	e.POST("/seasoning/add", handler.SeasoningAddPOST)
	e.POST("/recipe/add", handler.RecipeAddPOST)
	e.GET("/recipe/view/:id", handler.RecipeViewGET)
	e.POST("/recipe/info/add", handler.RecipeInfoAddPOST)
	e.GET("/recipe/info/view/:uid/:rid", handler.RecipeInfoViewGET)
	e.GET("/getImage/:get", handler.GetImage)
	e.POST("/room/add", handler.RoomAddPOST)
	e.GET("/roomId/duplicateCheck/:id", handler.RoomIDuplicationCheckGET)
	e.POST("/room/join", handler.JoinRoomPOST)
	e.POST("/machine/set", handler.SetMachinePOST)
	e.GET("/machine/get/:room_id/:recipe_id", handler.MachineInfoGET)
	e.GET("/machine/get/:room_id", handler.MachineRecipeNameGET)
	// e.GET("/image/up", handler.ImageUpload)

  e.Logger.Fatal(e.Start(":8081"))
}

// ハンドラーを定義
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "いちごたべさせろ〜")
}