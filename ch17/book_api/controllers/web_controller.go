package controllers

import (
	"book_api/database"
	"book_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ShowIndexPage(c *gin.Context) {
    var books []models.Book
    database.DB.Find(&books)


    c.HTML(http.StatusOK, "index.html", gin.H{
        "Books": books,
    })
}
