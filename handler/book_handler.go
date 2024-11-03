package handler

import (
	"github.com/gin-gonic/gin"
	"library/model"
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
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := bookHandler.bookService.AddBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}
func (bookHandler *BookHandler) UpdateBookHandler(c *gin.Context) {
	var pair struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&pair); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := bookHandler.bookService.UpdateBook(pair.ID, pair.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book updated"})
}
func (bookHandler *BookHandler) StoreStatsHandler(c *gin.Context) {
	bookCount, authorCount, err := bookHandler.bookService.GetStoreStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get store stats"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"book_count": bookCount, "distinct_authors": authorCount})
}
func (bookHandler *BookHandler) SearchBooksHandler(c *gin.Context) {
	title := c.Query("title")
	authorName := c.Query("author_name")
	priceRange := c.Query("price_range")
	books, err := bookHandler.bookService.SearchBooks(title, authorName, priceRange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search books"})
	}
	c.JSON(http.StatusOK, books)
}
