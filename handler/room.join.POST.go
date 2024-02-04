package handler

import (
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func JoinRoomPOST(c echo.Context) error {
	// リクエストのデータをモデルにバインド
	var joinRoom model.Room
    if err := c.Bind(&joinRoom); err != nil {
        c.Logger().Error(err)
        return err
  }

	room, err := model.JoinCheck(&joinRoom)
	if err != nil {
		// エラー時のハンドリング
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Internet Server Error"})
	}
	if room == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "Room not found"})
	}

	// パスワードが一致するかチェック
	err = model.CheckHashPassword(room.Password, joinRoom.Password)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "Password mismatch"})
	}

  return c.JSON(http.StatusOK, model.Response{Room: room})
}