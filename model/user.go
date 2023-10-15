package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       	int     `gorm:"primary_key;type:int"`
	Name 			string  `json:"name" gorm:"type:varchar(16);unique;not null"`
	Email 		string  `json:"emailRecipeName" gorm:"type:varchar(255);unique;not null"`
	Password 	string  `json:"password" gorm:"type:varchar(255);not null"`
	Icon			string	`json:"icon" gorm:"type:varchar(255);not null;default:'default-icon.png"`
}

func CreateUser(user *User) {
	db.Create(user)
}
func FindUser(u *User) User {
	var user User
	db.Where(u).First(&user)
	return user
}
