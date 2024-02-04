package model

import "gorm.io/gorm"

type SeasoningStatus struct {
	gorm.Model
	Status string `json:"status" gorm:"type:varchar(2);not null;unique"`
}

func CreateStatus(status *SeasoningStatus) {
	db.Create(status)
}

func FindSeasoningStatus(u *SeasoningStatus) SeasoningStatus {
	var status SeasoningStatus
	db.Where(u).First(&status)
	return status
}