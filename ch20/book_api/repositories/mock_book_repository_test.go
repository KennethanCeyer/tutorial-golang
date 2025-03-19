package repositories

import (
	"errors"
	"testing"

	"book_api/models"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMockBookRepository(t *testing.T) {
	// 1. Mock 데이터 초기화
	mockBooks := []models.Book{
		{Model: gorm.Model{ID: 1}, Title: "Mock 책 1", Author: "Mock 저자 1", Year: 2023},
		{Model: gorm.Model{ID: 2}, Title: "Mock 책 2", Author: "Mock 저자 2", Year: 2022},
	}

	// 2. MockBookRepository 생성
	repo := &MockBookRepository{MockBooks: mockBooks}

	// 3. FetchBooks 테스트
	books, err := repo.FetchBooks()
	assert.NoError(t, err, "FetchBooks에서 에러가 발생하지 않아야 합니다.")
	assert.Equal(t, len(mockBooks), len(books), "책의 개수가 일치해야 합니다.")

	// 4. FetchBookByID 테스트
	book, err := repo.FetchBookByID(1)
	assert.NoError(t, err, "FetchBookByID에서 에러가 발생하지 않아야 합니다.")
	assert.Equal(t, "Mock 책 1", book.Title, "책의 제목이 일치해야 합니다.")

	// 5. CreateBook 테스트
	newBook := models.Book{Model: gorm.Model{ID: 3}, Title: "새로운 책", Author: "새로운 저자", Year: 2024}
	err = repo.CreateBook(newBook)
	assert.NoError(t, err, "CreateBook에서 에러가 발생하지 않아야 합니다.")
	assert.Equal(t, 3, len(repo.MockBooks), "MockBooks의 길이가 3이어야 합니다.")

	// 6. UpdateBook 테스트
	updatedBook := models.Book{Model: gorm.Model{ID: 1}, Title: "수정된 책", Author: "수정된 저자", Year: 2025}
	err = repo.UpdateBook(updatedBook)
	assert.NoError(t, err, "UpdateBook에서 에러가 발생하지 않아야 합니다.")
	book, _ = repo.FetchBookByID(1)
	assert.Equal(t, "수정된 책", book.Title, "책 제목이 수정되지 않았습니다.")

	// 7. DeleteBook 테스트
	err = repo.DeleteBook(1)
	assert.NoError(t, err, "DeleteBook에서 에러가 발생하지 않아야 합니다.")
	assert.Equal(t, 2, len(repo.MockBooks), "MockBooks의 길이가 2여야 합니다.")

	// 8. MockErr 동작 테스트
	expectedErr := errors.New("mock error")
	repo.MockErr = expectedErr

	// 9. FetchBooks 에러 발생 확인
	_, err = repo.FetchBooks()
	assert.EqualError(t, err, expectedErr.Error(), "FetchBooks에서 MockErr가 발생해야 합니다.")

	// 10. FetchBookByID 에러 발생 확인
	_, err = repo.FetchBookByID(2)
	assert.EqualError(t, err, expectedErr.Error(), "FetchBookByID에서 MockErr가 발생해야 합니다.")

	// 11. CreateBook 에러 발생 확인
	err = repo.CreateBook(models.Book{Model: gorm.Model{ID: 4}, Title: "에러 책", Author: "에러 저자", Year: 2025})
	assert.EqualError(t, err, expectedErr.Error(), "CreateBook에서 MockErr가 발생해야 합니다.")

	// 12. UpdateBook 에러 발생 확인
	err = repo.UpdateBook(models.Book{Model: gorm.Model{ID: 2}, Title: "에러 수정", Author: "에러 저자", Year: 2026})
	assert.EqualError(t, err, expectedErr.Error(), "UpdateBook에서 MockErr가 발생해야 합니다.")

	// 13. DeleteBook 에러 발생 확인
	err = repo.DeleteBook(2)
	assert.EqualError(t, err, expectedErr.Error(), "DeleteBook에서 MockErr가 발생해야 합니다.")
}
