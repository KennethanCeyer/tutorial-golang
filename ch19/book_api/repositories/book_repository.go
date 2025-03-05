package repositories

import "book_api/models"

type BookRepository interface {
	FetchBooks() ([]models.Book, error)
	FetchBookByID(id uint) (models.Book, error)
	CreateBook(book models.Book) error
	UpdateBook(book models.Book) error
	DeleteBook(id uint) error
}
