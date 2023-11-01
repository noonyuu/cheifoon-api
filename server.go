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
  e.GET("/", hello)
	e.GET("/reset", handler.ResetGET)
	// e.POST("/user/add", handler.UserAddPOST)
	e.POST("/seasoning/add", handler.SeasoningAddPOST)
	e.POST("/recipe/add", handler.RecipeAddPOST)
	e.GET("/recipe/view/:id", handler.RecipeViewGET)
	e.POST("/menu/add", handler.MenuAddPOST)
	e.GET("/menu/view/:uid/:rid", handler.MenuViewGET)
	e.GET("/getImage/:get", handler.GetImage)
	// e.GET("/image/up", handler.ImageUpload)

  // サーバーをポート番号1323で起動
  e.Logger.Fatal(e.Start(":8081"))
}

// ハンドラーを定義
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

// package main

// import (
//     "fmt"
//     "net/http"
//     "github.com/labstack/echo"
//     "github.com/labstack/echo/middleware"
//     "os" // osパッケージを追加
//     "io" // ioパッケージを追加
// )

// func main() {
//     e := echo.New()
    
//     // ミドルウェアを設定（CORS対応など）
//     e.Use(middleware.CORS())
    
//     // 画像アップロードのハンドラ
//     e.POST("/upload-image", uploadImage)
    
//     e.Start(":8081")
// }

// func uploadImage(c echo.Context) error {
// 		fmt.Println("s")
//     // アップロードされた画像を取得
//     file, err := c.FormFile("image")
//     if err != nil {
//         return err
//     }
    
//     // アップロードされた画像を保存
//     src, err := file.Open()
//     if err != nil {
//         return err
//     }
//     defer src.Close()

//     // 保存先のファイルパスを指定
//     dst, err := os.Create("/public/uploads/" + file.Filename)
//     if err != nil {
//         return err
//     }
//     defer dst.Close()
    
//     // ファイルのコピー
//     if _, err = io.Copy(dst, src); err != nil {
//         return err
//     }

//     // ここでデータベースにファイルのパスやメタデータを保存する処理を実装
    
//     return c.String(http.StatusOK, "画像がアップロードされました")
// }
