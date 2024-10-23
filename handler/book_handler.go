package handler

import (
	"github.com/gin-gonic/gin"
	"library/service"
	"net/http"
)

type BookHandler struct {
	bookService *service.BookService
}

func NewBookHandler(bookService *service.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (bookHandler *BookHandler) GetBookHandler(c *gin.Context) {
	id := c.Query("id")
	book, err := bookHandler.bookService.GetBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}
func (bookHandler *BookHandler) DeleteBookHandler(c *gin.Context) {
	id := c.Query("id")
	err := bookHandler.bookService.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
func (bookHandler *BookHandler) AddBookHandler(c *gin.Context) {

}
func (bookHandler *BookHandler) UpdateBookHandler(c *gin.Context)  {}
func (bookHandler *BookHandler) SearchBooksHandler(c *gin.Context) {}
func (bookHandler *BookHandler) StoreStatsHandler(c *gin.Context)  {}
