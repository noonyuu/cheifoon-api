package model

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	ID       		int     `gorm:"primary_key;type:int"`
	UserId      int     `json:"user_id" gorm:"primary_key;type:int"`
	RecipeName 	string  `json:"recipe_name" gorm:"type:varchar(255);not null"`
	MenuImage 	string  `json:"menu_image" gorm:"type:varchar(255);not null"`
	User 				User 		`json:"users" gorm:"foreignKey:user_id"`
}

func CreateRecipe(recipe *Recipe) error {
	result := db.Create(recipe)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindRecipe(id int) ([]Recipe, error) {
    var recipe []Recipe
    err := db.Table("recipes").
        Select("id, user_id, recipe_name, menu_image").
        Where("user_id = ?", id).
        Scan(&recipe).Error
    if err != nil {
        return nil, err
    }
    return &recipe, nil
}