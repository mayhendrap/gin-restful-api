package db

import (
	"fmt"
	"github.com/mayhendrap/gin-restful-api/config"
	"github.com/mayhendrap/gin-restful-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	err = db.AutoMigrate(&models.Book{}, &models.User{})
	if err != nil {
		fmt.Println("Error migrating models")
		return
	}

	DB = db
}
