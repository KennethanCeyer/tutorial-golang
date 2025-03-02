package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"book_api/models"
	"book_api/repositories"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetBooks(t *testing.T) {
	// 1. Mock 데이터 정의
	mockBooks := []models.Book{
		{
			Model:  gorm.Model{ID: 1},
			Title:  "테스트 책 1",
			Author: "테스트 저자 1",
			Year:   2023,
		},
		{
			Model:  gorm.Model{ID: 2},
			Title:  "테스트 책 2",
			Author: "테스트 저자 2",
			Year:   2022,
		},
	}

	// 2. Mock Repository와 Controller 초기화
	mockRepo := &repositories.MockBookRepository{MockBooks: mockBooks}
	controller := &BookController{Repository: mockRepo}

	// 3. Gin Context Mocking
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// 4. 핸들러 실행
	controller.GetBooks(c)

	// 5. 결과 검증
	assert.Equal(t, http.StatusOK, w.Code)

	// 응답 데이터 비교
	var response []models.Book
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, len(mockBooks), len(response)) // 길이 비교

	for i, mockBook := range mockBooks {
		assert.Equal(t, mockBook.ID, response[i].ID, "ID가 일치하지 않습니다.")
		assert.Equal(t, mockBook.Title, response[i].Title, "Title이 일치하지 않습니다.")
		assert.Equal(t, mockBook.Author, response[i].Author, "Author가 일치하지 않습니다.")
		assert.Equal(t, mockBook.Year, response[i].Year, "Year가 일치하지 않습니다.")
	}
}

func TestDeleteBook(t *testing.T) {
    // 1. Mock 데이터 및 Repository 생성
    mockBooks := []models.Book{
        {Model: gorm.Model{ID: 1}, Title: "테스트 책", Author: "테스트 저자", Year: 2023},
    }
    mockRepo := &repositories.MockBookRepository{MockBooks: mockBooks}
    bookController := &BookController{Repository: mockRepo}

    // 2. Mock Gin Context 생성
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // 3. Mock Context에 필요한 데이터 설정
    c.Set("book", mockBooks[0]) // BookLoader 미들웨어의 동작을 대신 수행

    // 4. 핸들러 실행
    bookController.DeleteBook(c)

    // 5. 결과 검증
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "삭제되었습니다")
    assert.Equal(t, 0, len(mockRepo.MockBooks)) // MockBooks 리스트가 비었는지 확인
}

func TestGetBookByID_Success(t *testing.T) {
	// 1. Mock 데이터 정의
	mockBook := models.Book{Model: gorm.Model{ID: 1}, Title: "테스트 책", Author: "테스트 저자", Year: 2023}
	// 2. Mock Repository와 Controller 초기화
	mockRepo := &repositories.MockBookRepository{MockBooks: []models.Book{mockBook}}
	controller := &BookController{Repository: mockRepo}

	// 3. Gin Context 생성 및 'book' 데이터 설정 (BookLoader 역할 대체)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("book", mockBook)

	// 4. 핸들러 실행
	controller.GetBookByID(c)

	// 5. 결과 검증
	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Book
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, mockBook.ID, response.ID)
	assert.Equal(t, mockBook.Title, response.Title)
}

func TestGetBookByID_NoBook(t *testing.T) {
	// 1. Controller 초기화 (Mock Repository: 빈 데이터)
	controller := &BookController{Repository: &repositories.MockBookRepository{}}
	// 2. Gin Context 생성 (book 데이터 미설정)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// 3. 핸들러 실행
	controller.GetBookByID(c)

	// 4. 결과 검증
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "책 데이터를 로드하지 못했습니다")
}

func TestGetBookByID_InvalidType(t *testing.T) {
	// 1. Controller 초기화 (Mock Repository: 빈 데이터)
	controller := &BookController{Repository: &repositories.MockBookRepository{}}
	// 2. Gin Context 생성 및 'book' 데이터에 잘못된 타입 설정
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("book", "invalid type")

	// 3. 핸들러 실행
	controller.GetBookByID(c)

	// 4. 결과 검증
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "잘못된 책 데이터 형식")
}

func TestCreateBook(t *testing.T) {
	// 1. Mock Repository와 Controller 초기화 (빈 데이터)
	mockRepo := &repositories.MockBookRepository{MockBooks: []models.Book{}}
	controller := &BookController{Repository: mockRepo}

	// 2. 새 책 데이터 정의 및 JSON 직렬화
	newBook := models.Book{Title: "새 책", Author: "새 저자", Year: 2023}
	jsonData, err := json.Marshal(newBook)
	assert.NoError(t, err)

	// 3. Gin Context 생성 및 POST 요청 설정
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/books", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")

	// 4. 핸들러 실행
	controller.CreateBook(c)

	// 5. 결과 검증
	assert.Equal(t, http.StatusCreated, w.Code)
	var response models.Book
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, newBook.Title, response.Title)
	assert.Equal(t, newBook.Author, response.Author)
	assert.Equal(t, newBook.Year, response.Year)
	assert.Equal(t, 1, len(mockRepo.MockBooks))
}

func TestUpdateBook(t *testing.T) {
	// 1. 초기 데이터 정의 및 Mock Repository/Controller 초기화
	initialBook := models.Book{Model: gorm.Model{ID: 1}, Title: "원본 책", Author: "원본 저자", Year: 2020}
	mockRepo := &repositories.MockBookRepository{MockBooks: []models.Book{initialBook}}
	controller := &BookController{Repository: mockRepo}

	// 2. 업데이트할 데이터 정의 및 JSON 직렬화
	updatedPayload := models.Book{Title: "업데이트 책", Author: "업데이트 저자", Year: 2023}
	jsonData, err := json.Marshal(updatedPayload)
	assert.NoError(t, err)

	// 3. Gin Context 생성, PUT 요청 설정 및 'book' 데이터 설정 (BookLoader 역할 대체)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("PUT", "/books/1", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("book", initialBook)

	// 4. 핸들러 실행
	controller.UpdateBook(c)

	// 5. 결과 검증
	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Book
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, initialBook.ID, response.ID)
	assert.Equal(t, updatedPayload.Title, response.Title)
	assert.Equal(t, updatedPayload.Author, response.Author)
	assert.Equal(t, updatedPayload.Year, response.Year)
	assert.Equal(t, response, mockRepo.MockBooks[0])
}
