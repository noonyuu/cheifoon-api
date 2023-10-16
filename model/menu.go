package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	ID					int		 `gorm:"primary_key;type:int"`
	RecipeId		int		 `json:"recipe_id" gorm:"primary_key;type:int"`
	UserId      int    `json:"user_id" gorm:"primary_key;type:int"`
	SeasoningId	int		 `json:"seasoning_id":"type:int;not null"`
	TableSpoon	int		 `json:"table_spoon":"type:int;not null"`
	TeaSpoon		int		 `json:"tea_spoon":"type:int;not null"`
	Recipe			Recipe `json:"recipe" gorm:"foreignKey:recipe_id"`
	AdminSeasoning AdminSeasoning `json:"user_seasoning" gorm:"foreignKey:seasoning_id"`
	User 				User 		`json:"users" gorm:"foreignKey:user_id"`
}


func CreateMenu(menu *model.Menu) error {
    return db.Create(menu).Error
}

func FindMenu(user_id *int, recipe_id *int) ([]Menu, error) {
	var menus []Menu
	err := db.Table("menus").
			Select("id, recipe_id ,user_id, seasoning_id, table_spoon, tea_spoon").
			Where("user_id = ? AND recipe_id = ?", user_id, recipe_id).
			Scan(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}