package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Room struct {
    gorm.Model
    RoomID    string `json:"room_id" gorm:"type:varchar(16);not null"`
    RoomName  string `json:"room_name" gorm:"type:varchar(16);not null"`
    Password  string `json:"password" gorm:"type:varchar(255);not null"`
    Email     string `json:"email" gorm:"type:varchar(255);unique;not null"`
    Icon      string `json:"icon" gorm:"type:varchar(64);default:'default-icon.png'"`
}

func CreateRoom(room *Room) error {
	result := db.Create(room)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindRoom(r *Room) Room {
	var room Room
	db.Where(r).First(&room)
	return room
}

// ルームid重複チェック
func Duplicate(id *string) (bool, error) {
	var count int64
	err := db.Table("rooms").
		Select("room_id").
		Where("room_id = ?", id).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func JoinCheck(room *Room) (*Room, error) {
	var foundRoom Room
	err := db.Table("rooms").
		Where("room_id = ?", room.RoomID).
		First(&foundRoom).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // レコードが見つからなかった場合
		}
		return nil, err // それ以外のエラー
	}
	return &foundRoom, nil // レコードが見つかった場合
}

// 暗号化
func PasswordEncrypt(password string) (string, error) {
	// costはストレッチ回数を決めるパラメータ
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPassword), err
}

// 暗号化パスワードと比較
func CheckHashPassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

type Response struct {
    Room     *Room    `json:"room,omitempty"`
    Message  string   `json:"message,omitempty"`
}