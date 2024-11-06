package repository

import "library/model"

type BookRepository interface {
	GetBookByID(id string) (*model.Book, error)
	DeleteBookByID(id string) error
	UpdateBook(id string, title string) error
	AddBook(book *model.Book) (string, error)
	GetStoreStats() (int, int, error)
	SearchBooks(title, authorName, priceRange string) ([]*model.Book, error)
}
