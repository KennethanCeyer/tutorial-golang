package routes

import (
	"book_api/controllers"
	"book_api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, bookController *controllers.BookController) {
	r.LoadHTMLGlob("templates/*") // HTML 템플릿 로드

	// 기본 웹 페이지 라우트
	r.GET("/", bookController.ShowIndexPage)

	// API 라우트 그룹
	api := r.Group("/api")
	{
		api.GET("/books", bookController.GetBooks)
		api.GET("/books/:id", middlewares.BookLoader(bookController.Repository), bookController.GetBookByID)
		api.POST("/books", bookController.CreateBook)
		api.PUT("/books/:id", middlewares.BookLoader(bookController.Repository), bookController.UpdateBook)
		api.DELETE("/books/:id", middlewares.BookLoader(bookController.Repository), bookController.DeleteBook)
	}
}
