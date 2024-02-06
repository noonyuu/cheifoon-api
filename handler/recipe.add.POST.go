package handler

import (
    "myapp/model"
    "net/http"
    "fmt"

    "strconv"
    // "reflect"

    "github.com/labstack/echo/v4"
)

func RecipeAddPOST(c echo.Context) error {
	// 画像ファイルを取得
	file, err := c.FormFile("recipe_image")
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	// フォームデータからユーザーIDとレシピ名を取得
	roomIdStr := c.FormValue("room_id")
	roomId, err := strconv.Atoi(roomIdStr)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	recipeName := c.FormValue("recipe_name")

	// レシピデータを作成
	addRecipe := &model.Recipe{
		RoomID:     roomId,
		RecipeName: recipeName,
		RecipeImage:  file.Filename, // 仮のファイル名
	}

	fmt.Println(addRecipe.RoomID)
	// レシピ表のIDの最大値を取得
	maxID, err := model.MaxId()
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	// 新データ挿入時の仮ID
	tmpID := maxID + 1

	// RecipeImageのファイル名ようにintからstringへ変換
	roomID := strconv.FormatUint(uint64(addRecipe.RoomID), 10)
	// roomID := addRecipe.RoomID
  recipeID := strconv.FormatUint(uint64(tmpID), 10)

	// RecipeImageを退避
	// tmp := addRecipe.RecipeImage
  // RecipeImage を文字列として結合
	addRecipe.RecipeImage = roomID + recipeID + ".jpg"
	// レシピを追加
	err = model.CreateRecipe(addRecipe)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	if addRecipe.RecipeImage != "" {
		err := ImageUpload(c, file, addRecipe.RecipeImage)
		if err != nil {
			c.Logger().Error(err)
			return err
		}
	}

	newRecipeId := strconv.FormatUint(uint64(addRecipe.ID), 10)
	// JSON形式のレスポンスを生成
	response := map[string]interface{}{
			"id": newRecipeId,
	}

  return c.JSON(http.StatusCreated, response)

}



// func RecipeAddPOST(c echo.Context) error {
//     // レシピデータを取得
//     addRecipe := new(model.Recipe)
//     if err := c.Bind(addRecipe); err != nil {
//         c.Logger().Error(err)
//         return err
//     }
		
// 		fmt.Println("room_id",addRecipe.RoomID)

// 		var err error
//     err = model.CreateRecipe(addRecipe)
//     if err != nil {
//         c.Logger().Error(err)
//         return err
//     }

//     return c.JSON(http.StatusCreated, addRecipe.ID)
// }


// func RecipeAddPOST(c echo.Context) error {
	
// 	// files, err := c.FormFile("image")
//   //   if err != nil {
//   //       return err
//   //   }
// 	room, err := c.FormFile("room_id")
//     if err != nil {
//         return err
//     }
// 	fmt.Println("aiueo",room)
// 	// レシピデータを取得
// 	addRecipe := new(model.Recipe)
// 	if err := c.Bind(addRecipe); err != nil {
// 		c.Logger().Error(err)
// 		return err
// 	}

// 	// // レシピ表のIDの最大値を取得
// 	// maxID, err := model.MaxId(addRecipe.RoomID)
// 	// if err != nil {
// 	// 	c.Logger().Error(err)
// 	// 	return err
// 	// }
// 	// tmpID := *maxID + 1

// 	// roomID := strconv.FormatUint(uint64(addRecipe.RoomID), 10)
//   // recipeID := strconv.FormatUint(uint64(tmpID), 10)

// 	// // RecipeImageを退避
// 	// tmp := addRecipe.RecipeImage
//   // // RecipeImage を文字列として結合
// 	// addRecipe.RecipeImage = roomID + recipeID + ".jpg"
// 	// fmt.Println("name" , tmp)
// 	// // レシピを追加
// 	err = model.CreateRecipe(addRecipe)
// 	if err != nil {
// 		c.Logger().Error(err)
// 		return err
// 	}
// 	// fmt.Println("name" , reflect.TypeOf(tmp))
// 	// fmt.Println(addRecipe.ID)
// 	// if addRecipe.RecipeImage != "" {
// 	// 	err := ImageUpload(c, uint(addRecipe.ID), string(addRecipe.RecipeImage), uint(addRecipe.RoomID))
// 	// 	if err != nil {
// 	// 		c.Logger().Error(err)
// 	// 		return err
// 	// 	}
// 	// }

// 	return c.JSON(http.StatusCreated,addRecipe.ID)
// }