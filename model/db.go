package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func OpenDB() {
	var err error
	dsn := "root:root@tcp(db:3306)/cheifoon?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func Migrate() {
    db.AutoMigrate(&SeasoningStatus{},&Room{}, &Recipe{},&AdminSeasoning{},&UserSeasoning{},&RecipeInfo{},&Machine{})
    // db.Model(&Recipe{}).AddForeignKey("room_id", "rooms(room_id)", "RESTRICT", "RESTRICT")
}
		// ,&Recipe{},&RecipeInfo{},&User{},&Machine{}
	// &AdminSeasoning{},&UserSeasoning{},&Room{},&Recipe{},&RecipeInfo{}

func DeleteDB() {
	// テーブルを削除
	db.Migrator().DropTable(&UserSeasoning{})
	db.Migrator().DropTable(&RecipeInfo{})
	db.Migrator().DropTable(&AdminSeasoning{})
	db.Migrator().DropTable(&Recipe{})
	// db.Migrator().DropTable(&User{})
	// db.Migrator().DropTable(&Room{})
	db.Migrator().DropTable(&Machine{})
	db.Migrator().DropTable(&SeasoningStatus{})
}

func NewDBConnection() (*gorm.DB, error) {
	dsn := "root:root@tcp(db:3306)/cheifoon?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db, err
}

func GetDB() *gorm.DB {
    return db
}
