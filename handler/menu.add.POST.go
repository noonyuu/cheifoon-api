// package handler

// import (
// 	"myapp/model"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func MenuAddPOST(c echo.Context) error {

// 	// 調味料データを取得
// 	addMenu := new(model.Menu)
// 	if err := c.Bind(addMenu); err != nil {
// 		c.Logger().Error(err)
// 		return err
// 	}

// 	// 調味料を追加
// 	model.CreateMenu(addMenu)

// 	return c.JSON(http.StatusCreated, addMenu)
// }

package handler

import (
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MenuAddPOST(c echo.Context) error {
	var addMenus []model.Menu
	if err := c.Bind(&addMenus); err != nil {
		c.Logger().Error(err)
		return err
	}

	// 複数のメニューデータをデータベースに追加
	createdMenus := make([]model.Menu, 0, len(addMenus))
	for _, addMenu := range addMenus {
		// 各メニューごとにデータベースに追加
		model.CreateMenu(&addMenu)
		createdMenus = append(createdMenus, addMenu)
	}

	return c.JSON(http.StatusCreated, createdMenus)
}
