package repositories

import (
	"book_api/models"
	"fmt"
)

type MockBookRepository struct {
	MockBooks []models.Book
}

func (r *MockBookRepository) FetchBooks() ([]models.Book, error) {
	return r.MockBooks, nil
}

func (r *MockBookRepository) FetchBookByID(id uint) (models.Book, error) {
	for _, book := range r.MockBooks {
		if book.ID == id {
			return book, nil
		}
	}
	return models.Book{}, fmt.Errorf("책을 찾을 수 없습니다")
}

func (r *MockBookRepository) CreateBook(book models.Book) error {
	r.MockBooks = append(r.MockBooks, book)
	return nil
}

func (r *MockBookRepository) UpdateBook(book models.Book) error {
	for i, b := range r.MockBooks {
		if b.ID == book.ID {
			r.MockBooks[i] = book
			return nil
		}
	}
	return nil
}

func (r *MockBookRepository) DeleteBook(id uint) error {
	for i, b := range r.MockBooks {
		if b.ID == id {
			r.MockBooks = append(r.MockBooks[:i], r.MockBooks[i+1:]...)
			return nil
		}
	}
	return nil
}
