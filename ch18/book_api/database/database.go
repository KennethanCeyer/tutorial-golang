package database

import (
	"book_api/models"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("데이터베이스 연결 실패:", err)
	}

	// 마이그레이션 수행
	DB.AutoMigrate(&models.Book{})
	log.Println("데이터베이스 초기화 완료")
}
