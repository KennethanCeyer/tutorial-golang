package routes

import (
	"book_api/controllers"

	"github.com/gin-gonic/gin"
)


func SetupRoutes(r *gin.Engine) {
    r.LoadHTMLGlob("templates/*")


    r.GET("/", controllers.ShowIndexPage)


    api := r.Group("/api")
    {
        api.GET("/books", controllers.GetBooks)
        api.GET("/books/:id", controllers.GetBookByID)
        api.POST("/books", controllers.CreateBook)
        api.PUT("/books/:id", controllers.UpdateBook)
        api.DELETE("/books/:id", controllers.DeleteBook)
    }
}
