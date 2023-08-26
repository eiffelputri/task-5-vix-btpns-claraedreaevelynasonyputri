package database

import (
	"task-5-vix-btpns-ClaraEdreaEvelynaSonyPutri/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// InitDB initializes the database connection
func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "btpns-photo.db")
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Photo{})

	return db
}
