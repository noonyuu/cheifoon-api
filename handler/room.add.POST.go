package handler

import (
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RoomAddPOST(c echo.Context) error {
	// リクエストのデータをモデルにバインド
	room := new(model.Room)
	if err := c.Bind(room); err != nil {
		c.Logger().Error(err)
		return err
	}

	// パスワードをハッシュ化
	encryptPass, err := model.PasswordEncrypt(room.Password)
	if err != nil {
		return err
	}

	// ハッシュ化したパスワードを設定
	room.Password = encryptPass

	// ルーム登録
	err = model.CreateRoom(room)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "room creation failed",
		}
	}


	return c.JSON(http.StatusCreated, room)
}
