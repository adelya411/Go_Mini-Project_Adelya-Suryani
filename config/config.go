package config

import (
	"tulispuisi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDataBase() {
	dsn := "root:@tcp(127.0.0.1:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Can't Connect to Database")
	}

	Migration()
}

func Migration() {
	DB.AutoMigrate(&models.Puisi{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Kategori{})
	DB.AutoMigrate(&models.Rating{})
}
