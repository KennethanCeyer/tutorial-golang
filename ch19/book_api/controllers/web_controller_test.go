package controllers

import (
	"errors"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"book_api/models"
	"book_api/repositories"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestShowIndexPage_Success(t *testing.T) {
	// 1. 테스트용 Mock 데이터 정의
	mockBooks := []models.Book{
		{Model: gorm.Model{ID: 1}, Title: "테스트 책 1", Author: "테스트 저자 1", Year: 2023},
		{Model: gorm.Model{ID: 2}, Title: "테스트 책 2", Author: "테스트 저자 2", Year: 2022},
	}

	// 2. Mock Repository와 Controller 초기화
	mockRepo := &repositories.MockBookRepository{MockBooks: mockBooks}
	webController := &WebController{Repository: mockRepo}

	// 3. 테스트용 컨텍스트 생성
	w := httptest.NewRecorder()
	c, engine := gin.CreateTestContext(w)

	// 4. HTML 템플릿 설정  
	gin.SetMode(gin.TestMode)
	tmpl := template.Must(template.New("tmpl").Parse(`
{{define "index.html"}}Index Page: {{len .Books}}{{end}}
{{define "error.html"}}Error Page: {{.error}}{{end}}
`))
	engine.SetHTMLTemplate(tmpl)

	// 5. 핸들러 실행
	webController.ShowIndexPage(c)

	// 6. 결과 검증
	assert.Equal(t, http.StatusOK, w.Code)
	expected := "Index Page: 2"
	assert.Contains(t, w.Body.String(), expected)
}

func TestShowIndexPage_Error(t *testing.T) {
	// 1. 테스트용 Mock 데이터 정의
	mockRepo := &repositories.MockBookRepository{
		MockErr: errors.New("fetch error"),
	}

	// 2. Controller 초기화
	webController := &WebController{Repository: mockRepo}

	// 3. 테스트용 컨텍스트 생성
	w := httptest.NewRecorder()
	c, engine := gin.CreateTestContext(w)

	// 4. HTML 템플릿 설정
	gin.SetMode(gin.TestMode)
	tmpl := template.Must(template.New("tmpl").Parse(`
{{define "index.html"}}Index Page: {{len .Books}}{{end}}
{{define "error.html"}}Error Page: {{.error}}{{end}}
`))
	engine.SetHTMLTemplate(tmpl)

	// 5. 핸들러 실행
	webController.ShowIndexPage(c)

	// 6. 결과 검증
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	expected := "Error Page: 데이터를 불러올 수 없습니다."
	assert.Contains(t, w.Body.String(), expected)
}
