package controllers

import (
	"book_api/models"
	"book_api/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	Repository repositories.BookRepository
}

func (bc *BookController) ShowIndexPage(c *gin.Context) {
	books, err := bc.Repository.FetchBooks()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "데이터를 불러올 수 없습니다."})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"Books": books})
}

func (bc *BookController) GetBooks(c *gin.Context) {
	books, err := bc.Repository.FetchBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "데이터를 가져오는 데 실패했습니다"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (bc *BookController) GetBookByID(c *gin.Context) {
	book, exists := c.Get("book")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "책 데이터를 로드하지 못했습니다"})
		return
	}

	bookModel, ok := book.(models.Book)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "잘못된 책 데이터 형식"})
		return
	}

	c.JSON(http.StatusOK, bookModel)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	book, exists := c.Get("book")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "책 데이터를 로드하지 못했습니다"})
		return
	}

	bookModel, ok := book.(models.Book)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "잘못된 책 데이터 형식"})
		return
	}

	if err := bc.Repository.DeleteBook(bookModel.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "책 삭제에 실패했습니다"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "책이 삭제되었습니다."})
}

func (bc *BookController) CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청 데이터"})
		return
	}
	if err := bc.Repository.CreateBook(newBook); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "책 생성에 실패했습니다"})
		return
	}
	c.JSON(http.StatusCreated, newBook)
}

func (bc *BookController) UpdateBook(c *gin.Context) {
	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청 데이터"})
		return
	}
	book, exists := c.Get("book")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "책 데이터를 로드하지 못했습니다"})
		return
	}
	bookModel, ok := book.(models.Book)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "잘못된 책 데이터 형식"})
		return
	}
	updatedBook.ID = bookModel.ID
	if err := bc.Repository.UpdateBook(updatedBook); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "책 수정에 실패했습니다"})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}
