package database

import (
	"book_api/models"

	"gorm.io/gorm"
)

func InitDB(dialector gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Book{}); err != nil {
		return nil, err
	}

	return db, nil
}
