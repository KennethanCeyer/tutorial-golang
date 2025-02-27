package database


import (
    "book_api/models"
    "github.com/glebarez/sqlite"
    "testing"
)


func TestInitDBWithMemory(t *testing.T) {
    // 메모리 기반 데이터베이스를 사용한 InitDB 호출
    db, err := InitDB(sqlite.Open(":memory:"))
    if err != nil {
        t.Fatalf("InitDB 실패: %v", err)
    }


    // 데이터 삽입 테스트
    testBook := models.Book{Title: "테스트 책", Author: "테스터", Year: 2023}
    result := db.Create(&testBook)
    if result.Error != nil {
        t.Fatalf("데이터 삽입 실패: %v", result.Error)
    }


    // 데이터 조회 테스트
    var retrievedBook models.Book
    if err := db.First(&retrievedBook, testBook.ID).Error; err != nil {
        t.Fatalf("데이터 조회 실패: %v", err)
    }


    // 데이터 검증
    if retrievedBook.Title != testBook.Title {
        t.Errorf("저장된 데이터와 조회된 데이터가 일치하지 않습니다. (저장: %v, 조회: %v)", testBook.Title, retrievedBook.Title)
    }
}
