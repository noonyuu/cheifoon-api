package model

import (
	"errors"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	UserId     int   `json:"user_id" gorm:"not null"`
	RecipeName string `json:"recipe_name" gorm:"type:varchar(255);not null"`
	MenuImage  string `json:"menu_image" gorm:"type:varchar(255);not null"`
	User       User   `json:"users" gorm:"foreignKey:user_id"`
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
	return recipe, nil
}

func MaxId(user_id int) (*uint, error) {
	var maxId uint
	err := db.Table("recipes").
		Select("COALESCE(MAX(id), 0)").
		Where("user_id = ?", user_id).
		Scan(&maxId).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &maxId, nil
}
