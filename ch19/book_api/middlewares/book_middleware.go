package middlewares

import (
	"book_api/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BookLoader(repo repositories.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 64)
		fmt.Printf("id=%v\n", id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 ID"})
			c.Abort()
			return
		}

		book, err := repo.FetchBookByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "책을 찾을 수 없습니다"})
			c.Abort()
			return
		}

		if book.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "책을 찾을 수 없습니다"})
			c.Abort()
			return
		}

		c.Set("book", book)
		c.Next()
	}
}
