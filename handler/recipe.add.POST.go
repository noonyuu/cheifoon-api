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
	file, err := c.FormFile("menu_image")
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	// フォームデータからユーザーIDとレシピ名を取得
	userIdStr := c.FormValue("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	recipeName := c.FormValue("recipe_name")

	// レシピデータを作成
	addRecipe := &model.Recipe{
		UserId:     userId,
		RecipeName: recipeName,
		MenuImage:  file.Filename, // 仮のファイル名
	}

	fmt.Println(addRecipe.UserId)
	// レシピ表のIDの最大値を取得
	maxID, err := model.MaxId(addRecipe.UserId)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	// 新データ挿入時の仮ID
	tmpID := *maxID + 1

	// menuImageのファイル名ようにintからstringへ変換
	userID := strconv.FormatUint(uint64(addRecipe.UserId), 10)
  recipeID := strconv.FormatUint(uint64(tmpID), 10)

	// menuImageを退避
	// tmp := addRecipe.MenuImage
  // MenuImage を文字列として結合
	addRecipe.MenuImage = userID + recipeID + ".jpg"
	// レシピを追加
	err = model.CreateRecipe(addRecipe)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	if addRecipe.MenuImage != "" {
		err := ImageUpload(c, file, addRecipe.MenuImage)
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
		
// 		fmt.Println("user_id",addRecipe.UserId)

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
// 	user, err := c.FormFile("user_id")
//     if err != nil {
//         return err
//     }
// 	fmt.Println("aiueo",user)
// 	// レシピデータを取得
// 	addRecipe := new(model.Recipe)
// 	if err := c.Bind(addRecipe); err != nil {
// 		c.Logger().Error(err)
// 		return err
// 	}

// 	// // レシピ表のIDの最大値を取得
// 	// maxID, err := model.MaxId(addRecipe.UserId)
// 	// if err != nil {
// 	// 	c.Logger().Error(err)
// 	// 	return err
// 	// }
// 	// tmpID := *maxID + 1

// 	// userID := strconv.FormatUint(uint64(addRecipe.UserId), 10)
//   // recipeID := strconv.FormatUint(uint64(tmpID), 10)

// 	// // menuImageを退避
// 	// tmp := addRecipe.MenuImage
//   // // MenuImage を文字列として結合
// 	// addRecipe.MenuImage = userID + recipeID + ".jpg"
// 	// fmt.Println("name" , tmp)
// 	// // レシピを追加
// 	err = model.CreateRecipe(addRecipe)
// 	if err != nil {
// 		c.Logger().Error(err)
// 		return err
// 	}
// 	// fmt.Println("name" , reflect.TypeOf(tmp))
// 	// fmt.Println(addRecipe.ID)
// 	// if addRecipe.MenuImage != "" {
// 	// 	err := ImageUpload(c, uint(addRecipe.ID), string(addRecipe.MenuImage), uint(addRecipe.UserId))
// 	// 	if err != nil {
// 	// 		c.Logger().Error(err)
// 	// 		return err
// 	// 	}
// 	// }

// 	return c.JSON(http.StatusCreated,addRecipe.ID)
// }