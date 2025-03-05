package repositories

import (
	"book_api/models"

	"gorm.io/gorm"
)

type GormBookRepository struct {
	DB *gorm.DB
}

func (r *GormBookRepository) FetchBooks() ([]models.Book, error) {
	var books []models.Book
	result := r.DB.Find(&books)
	return books, result.Error
}

func (r *GormBookRepository) FetchBookByID(id uint) (models.Book, error) {
	var book models.Book
	result := r.DB.First(&book, id)
	return book, result.Error
}

func (r *GormBookRepository) CreateBook(book models.Book) error {
	result := r.DB.Create(&book)
	return result.Error
}

func (r *GormBookRepository) UpdateBook(book models.Book) error {
	result := r.DB.Save(&book)
	return result.Error
}

func (r *GormBookRepository) DeleteBook(id uint) error {
	result := r.DB.Delete(&models.Book{}, id)
	return result.Error
}
