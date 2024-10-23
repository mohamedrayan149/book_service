package service

import (
	"library/model"
	"library/repository"
)

type BookService struct {
	bookRepository *repository.BookRepository
}

func NewBookService(bookRepository *repository.BookRepository) *BookService {
	return &BookService{bookRepository: bookRepository}
}

func (bookService *BookService) GetBook(bookId string) (*model.Book, error) {
	return bookService.bookRepository.GetBookByID(bookId)
}
func (bookService *BookService) DeleteBook(bookId string) error {
	return bookService.bookRepository.DeleteBookByID(bookId)
}
