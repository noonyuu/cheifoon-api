package model

import (
	"errors"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	RecipeName  string `json:"recipe_name" gorm:"type:varchar(255);not null"`
	RoomID      int `json:"room_id" gorm:"type:int;not null"`
	// Room        Room   `gorm:"foreignKey:RoomID"`
	RecipeImage string `json:"recipe_image" gorm:"type:varchar(255);not null"`
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
	// nub := "cheifoon"
	err := db.Table("recipes").
		Select("id, room_id, recipe_name, recipe_image").
		Where("room_id = ?", id).
		Scan(&recipe).Error
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func MaxId() (int, error) {
	var maxId int
	err := db.Table("recipes").
		Select("COALESCE(COUNT(id), 0)").
		// Where("room_id = ?", room_id).
		Scan(&maxId).
		Error
	if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return 0, nil
        }
        return 0, err
    }
    return maxId, nil
}
