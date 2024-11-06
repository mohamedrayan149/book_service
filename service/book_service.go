package service

import (
	"library/model"
	"library/repository"
)

type BookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) *BookService {
	return &BookService{bookRepository: bookRepository}
}

func (bookService *BookService) GetBookByID(bookId string) (*model.Book, error) {
	return bookService.bookRepository.GetBookByID(bookId)
}

func (bookService *BookService) DeleteBookByID(bookId string) error {
	return bookService.bookRepository.DeleteBookByID(bookId)
}

func (bookService *BookService) UpdateBook(bookId string, title string) error {
	return bookService.bookRepository.UpdateBook(bookId, title)
}

func (bookService *BookService) AddBook(book *model.Book) (string, error) {
	return bookService.bookRepository.AddBook(book)
}
func (bookService *BookService) GetStoreStats() (int, int, error) {
	return bookService.bookRepository.GetStoreStats()
}

func (bookService *BookService) SearchBooks(title, authorName, priceRange string) ([]*model.Book, error) {
	return bookService.bookRepository.SearchBooks(title, authorName, priceRange)
}
