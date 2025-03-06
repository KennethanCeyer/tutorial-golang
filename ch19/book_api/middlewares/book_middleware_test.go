package middlewares

import (
	"book_api/models"
	"book_api/repositories"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// 정상 케이스: 유효한 id로 요청 시 컨텍스트에 'book' 데이터가 저장되어야 한다.
func TestBookLoader(t *testing.T) {
	mockBooks := []models.Book{
		{Model: gorm.Model{ID: 1}, Title: "테스트 책", Author: "테스트 저자", Year: 2023},
	}
	mockRepo := &repositories.MockBookRepository{MockBooks: mockBooks}

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/books/:id", BookLoader(mockRepo), func(c *gin.Context) {
		book, exists := c.Get("book")
		assert.True(t, exists, "Context에 'book' 데이터가 저장되어야 합니다.")

		bookModel, ok := book.(models.Book)
		assert.True(t, ok, "Context의 'book' 데이터가 models.Book 타입이어야 합니다.")

		c.JSON(http.StatusOK, bookModel)
	})

	req := httptest.NewRequest("GET", "/books/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "테스트 책")
}

// 잘못된 ID 케이스: 숫자가 아닌 id를 전달하면 400 Bad Request가 반환되어야 한다.
func TestBookLoader_InvalidID(t *testing.T) {
	mockRepo := &repositories.MockBookRepository{}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/books/:id", BookLoader(mockRepo), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "should not reach"})
	})

	req := httptest.NewRequest("GET", "/books/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "유효하지 않은 ID")
}

// 존재하지 않는 책: Repository에 해당 id의 책이 없으면 404 Not Found가 반환되어야 한다.
func TestBookLoader_NotFound(t *testing.T) {
	// 빈 MockBookRepository 사용: id=1인 책이 없음.
	mockRepo := &repositories.MockBookRepository{MockBooks: []models.Book{}}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/books/:id", BookLoader(mockRepo), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "should not reach"})
	})

	req := httptest.NewRequest("GET", "/books/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "책을 찾을 수 없습니다")
}
