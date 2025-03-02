package repositories

import (
	"testing"

	"book_api/database"
	"book_api/models"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGormBookRepository(t *testing.T) {
	// SQLite 메모리 데이터베이스 초기화
	db, err := database.InitDB(sqlite.Open(":memory:"))
	assert.NoError(t, err, "데이터베이스 연결 실패")

	// 테이블 생성
	err = db.AutoMigrate(&models.Book{})
	assert.NoError(t, err, "테이블 생성 실패")

	// GormBookRepository 초기화
	repo := &GormBookRepository{DB: db}

	// CreateBook 테스트
	book := models.Book{Title: "테스트 책", Author: "테스트 저자", Year: 2023}
	err = repo.CreateBook(book)
	assert.NoError(t, err, "CreateBook에서 에러가 발생하지 않아야 합니다.")

	// FetchBooks 테스트
	books, err := repo.FetchBooks()
	assert.NoError(t, err, "FetchBooks에서 에러가 발생하지 않아야 합니다.")
	assert.Equal(t, 1, len(books), "책의 개수가 1이어야 합니다.")

	// FetchBookByID 테스트
	fetchedBook, err := repo.FetchBookByID(1)
	assert.NoError(t, err, "FetchBookByID에서 에러가 발생하지 않아야 합니다.")
	assert.Equal(t, "테스트 책", fetchedBook.Title, "책의 제목이 일치해야 합니다.")

	// UpdateBook 테스트
	updatedBook := models.Book{
		Model:  gorm.Model{ID: 1},
		Title:  "업데이트된 책",
		Author: "업데이트된 저자",
		Year:   2024,
	}
	err = repo.UpdateBook(updatedBook)
	assert.NoError(t, err, "UpdateBook에서 에러가 발생하지 않아야 합니다.")

	// DeleteBook 테스트
	err = repo.DeleteBook(1)
	assert.NoError(t, err, "DeleteBook에서 에러가 발생하지 않아야 합니다.")
	books, err = repo.FetchBooks()
	assert.NoError(t, err, "FetchBooks에서 에러가 발생하지 않아야 합니다.")
	assert.Equal(t, 0, len(books), "책의 개수가 0이어야 합니다.")
}
