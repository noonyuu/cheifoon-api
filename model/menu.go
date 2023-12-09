package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	RecipeId     int            `json:"recipe_id" gorm:"type:int;not null"`
	UserId       int            `json:"user_id" gorm:"type:int;not null"`
	SeasoningId  int            `json:"seasoning_id" gorm:"type:int;not null"`
	TableSpoon   int            `json:"table_spoon" gorm:"type:int;not null"`
	TeaSpoon     int            `json:"tea_spoon" gorm:"type:int;not null"`
	Recipe       Recipe         `json:"recipe" gorm:"foreignKey:RecipeId"`
	AdminSeasoning AdminSeasoning `json:"user_seasoning" gorm:"foreignKey:SeasoningId"`
	User         User           `json:"users" gorm:"foreignKey:UserId"`
	RecipeName   string         `json:"recipe_name" gorm:"-"`
	SeasoningName string				`json:"seasoning_name" gorm:"-"`
}

func CreateMenu(menu *Menu) {
	db.Create(menu)
}

func FindMenu(user_id *int, recipe_id *int) ([]Menu, error) {
	var menus []Menu
	err := db.Table("menus").
		Select("menus.id, menus.recipe_id, menus.user_id, menus.seasoning_id, menus.table_spoon, menus.tea_spoon").
		Joins("LEFT JOIN recipes ON menus.recipe_id = recipes.id").
		Joins("JOIN admin_seasonings ON menus.seasoning_id = admin_seasonings.id").
		Where("menus.user_id = ? AND menus.recipe_id = ?", user_id, recipe_id).
		Scan(&menus).Error
	if err != nil {
		return nil, err
	}

	// 結合後のデータを手動で取得して設定
	for i := range menus {
		var recipeName string
		var seasoningName string
		row_recipe := db.Table("recipes").
			Select("recipe_name").
			Where("id = ?", menus[i].RecipeId).
			Row()
		if err := row_recipe.Scan(&recipeName); err != nil {
			return nil, err
		}

		row_seasoning := db.Table("admin_seasonings").
			Select("seasoning_name").
			Where("id = ?", menus[i].SeasoningId).
			Row()
		if err = row_seasoning.Scan(&seasoningName); err != nil {
			return nil,err
		}

		menus[i].RecipeName = recipeName
		menus[i].SeasoningName = seasoningName
	}

	return menus, nil
}
