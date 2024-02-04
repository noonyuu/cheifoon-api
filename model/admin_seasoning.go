package model

import "gorm.io/gorm"

type AdminSeasoning struct {
	gorm.Model
	ID            int             `gorm:"primary_key;type:int"`
	SeasoningName string          `json:"seasoning_name" gorm:"type:varchar(255);unique;not null"`
	TeaSecond     float64         `json:"tea_second" gorm:"type:double;not null"`
	StatusId      int             `json:"status_id" gorm:"type:int;not null"`
	SeasoningStatus SeasoningStatus `json:"-" gorm:"foreignKey:StatusId"`
	BottleImage   string          `json:"bottle_image" gorm:"type:varchar(255);not null"`
}

func CreateSeasoning(admin *AdminSeasoning) {
	db.Create(admin)
}

func FindAdmin(a *AdminSeasoning) AdminSeasoning {
	var admin AdminSeasoning
	db.Where(a).First(&admin)
	return admin
}
