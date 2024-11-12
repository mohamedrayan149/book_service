package datastore

import (
	"library/config"
	"library/model"
)

type BookStore interface {
	GetBook(id string) (*model.Book, error)
	DeleteBook(id string) error
	UpdateBook(id string, title string) error
	AddBook(book *config.BookAddData) (string, error)
	GetStoreStats() (int, int, error)
	SearchBooks(title, authorName, priceRange string) ([]*model.Book, error)
}

type UserActivity interface {
	LogUserAction(username, action string) error
	GetLastUserActions(username string) ([]string, error)
}
