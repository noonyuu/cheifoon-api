package model

import "gorm.io/gorm"

type RecipeInfo struct {
	gorm.Model
	RecipeId      int            `json:"recipe_id" gorm:"type:int;not null"`
	Recipe        Recipe         `json:"recipe" gorm:"foreignKey:RecipeId"`
	RoomID        int         `json:"room_id" gorm:"type:int;not null"`
	SeasoningId   int            `json:"seasoning_id" gorm:"type:int;not null"`
	TableSpoon    int            `json:"table_spoon" gorm:"type:int;not null"`
	TeaSpoon      int            `json:"tea_spoon" gorm:"type:int;not null"`
	AdminSeasoning AdminSeasoning `json:"user_seasoning" gorm:"foreignKey:SeasoningId"`
	// Room					Room						`json:"room" gorm:"foreignKey:RoomID"`
	RecipeName    string         `json:"recipe_name" gorm:"-"`
	SeasoningName string         `json:"seasoning_name" gorm:"-"`
}

func CreateRecipeInfo(recipeInfo *RecipeInfo) {
	db.Create(recipeInfo)
}

func FindRecipeInfo(roomId *int, recipe_id *int) ([]RecipeInfo, error) {
    var recipeInfos []RecipeInfo
    err := db.Table("recipe_infos").
			Select("recipe_infos.id, recipe_infos.recipe_id, recipe_infos.room_id, recipe_infos.seasoning_id, recipe_infos.table_spoon, recipe_infos.tea_spoon, recipes.recipe_name, admin_seasonings.seasoning_name").
			Joins("LEFT JOIN recipes ON recipe_infos.recipe_id = recipes.id").
			Joins("JOIN admin_seasonings ON recipe_infos.seasoning_id = admin_seasonings.id").
			Where("recipe_infos.room_id = ? AND recipe_infos.recipe_id = ?", roomId, recipe_id).
			Scan(&recipeInfos).Error
	if err != nil {
		return nil, err
	}

	// 結合後のデータを手動で取得して設定
	for i := range recipeInfos {
		var recipeName string
		var seasoningName string
		row_recipe := db.Table("recipes").
			Select("recipe_name").
			Where("id = ?", recipeInfos[i].RecipeId).
			Row()
		if err := row_recipe.Scan(&recipeName); err != nil {
			return nil, err
		}

		row_seasoning := db.Table("admin_seasonings").
			Select("seasoning_name").
			Where("id = ?", recipeInfos[i].SeasoningId).
			Row()
		if err = row_seasoning.Scan(&seasoningName); err != nil {
			return nil,err
		}

		recipeInfos[i].RecipeName = recipeName
		recipeInfos[i].SeasoningName = seasoningName
	}

	return recipeInfos, nil
}