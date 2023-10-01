package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_workshop_db"))
	if err != nil {
		panic(err)
	}

	product := Product{}
	database.AutoMigrate(&product)

	DB = database
}
