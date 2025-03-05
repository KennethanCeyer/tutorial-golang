package controllers

import (
	"book_api/database"
	"book_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetBooks(c *gin.Context) {
    var books []models.Book
    result := database.DB.Find(&books)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, books)
}


func GetBookByID(c *gin.Context) {
    var book models.Book
    if err := database.DB.First(&book, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "책을 찾을 수 없습니다."})
        return
    }
    c.JSON(http.StatusOK, book)
}


func CreateBook(c *gin.Context) {
    var newBook models.Book
    if err := c.ShouldBindJSON(&newBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }


    result := database.DB.Create(&newBook)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusCreated, newBook)
}


func UpdateBook(c *gin.Context) {
    var book models.Book
    if err := database.DB.First(&book, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "책을 찾을 수 없습니다."})
        return
    }


    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }


    result := database.DB.Save(&book)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, book)
}


func DeleteBook(c *gin.Context) {
    var book models.Book
    if err := database.DB.First(&book, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "책을 찾을 수 없습니다."})
        return
    }


    result := database.DB.Delete(&book)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "책이 삭제되었습니다."})
}
