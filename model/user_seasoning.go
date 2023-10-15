package model

import "gorm.io/gorm"

type UserSeasoning struct {
	gorm.Model
	ID       			int     `gorm:"primary_key;type:int"`
	SeasoningId 	int  `json:"seasoning_id" gorm:"type:int;unique;not null"`
	AdminSeasoning AdminSeasoning `json:"admin_seasoning" gorm:"foreignKey:SeasoningId"`
}

func AddSeasoning(user *UserSeasoning) {
	db.Create(user)
}

func FindUserSeasoning(u *UserSeasoning) UserSeasoning {
	var user UserSeasoning
	db.Where(u).First(&user)
	return user
}
