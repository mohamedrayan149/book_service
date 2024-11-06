package handler

import (
	"github.com/gin-gonic/gin"
	"library/service"
	"library/utilites"
	"net/http"
)

// Constants for parameter names
const (
	IDParam         = "id"
	TitleParam      = "title"
	AuthorParam     = "author_name"
	PriceRangeParam = "price_range"
	EmptyString     = ""
)

// Constants for error messages
const (
	ErrFailedToGetBook     = "Failed to get book"
	ErrFailedToAddBook     = "Failed to add book"
	ErrFailedToUpdateBook  = "Failed to update book"
	ErrFailedToDeleteBook  = "Failed to delete book"
	ErrFailedToGetStats    = "Failed to get store stats"
	ErrFailedToSearchBooks = "Failed to search books"
	ErrParameterRequired   = "At least one of the following parameters is required: title, author_name, or price_range"
	ErrInvalidPriceRange   = "price_range must be in the format 'min-max' with numeric values"
)

// Constants for success messages
const (
	SuccessBookDeleted = "Book deleted successfully"
	SuccessBookUpdated = "Book updated successfully"
)

// Constants for response fields
const (
	BookCountField       = "book_count"
	DistinctAuthorsField = "distinct_authors"
	MessageField         = "message"
	ErrorField           = "error"
)

type BookHandler struct {
	bookService *service.BookService
}

func NewBookHandler(bookService *service.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (bookHandler *BookHandler) GetBookHandler(c *gin.Context) {
	id, ok := utilites.GetIDFromQuery(c)
	if !ok {
		return // The error response is already handled by GetIDFromQuery
	}
	_, ok = utilites.GetUsernameFromQuery(c)
	if !ok {
		return
	}
	book, err := bookHandler.bookService.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{ErrorField: ErrFailedToGetBook})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (bookHandler *BookHandler) DeleteBookHandler(c *gin.Context) {
	id, ok := utilites.GetIDFromQuery(c)
	if !ok {
		return // The error response is already handled by GetIDFromQuery
	}
	_, ok = utilites.GetUsernameFromQuery(c)
	if !ok {
		return
	}
	err := bookHandler.bookService.DeleteBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{ErrorField: ErrFailedToDeleteBook})
		return
	}
	c.JSON(http.StatusOK, gin.H{MessageField: SuccessBookDeleted})
}

func (bookHandler *BookHandler) AddBookHandler(c *gin.Context) {
	book, ok := utilites.BindBookData(c)
	if !ok {
		return
	}
	id, err := bookHandler.bookService.AddBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ErrorField: ErrFailedToAddBook})
		return
	}
	c.JSON(http.StatusCreated, gin.H{IDParam: id})
}

func (bookHandler *BookHandler) UpdateBookHandler(c *gin.Context) {
	data, ok := utilites.BindBookUpdateData(c)
	if !ok {
		return
	}
	err := bookHandler.bookService.UpdateBook(data.ID, data.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ErrorField: ErrFailedToUpdateBook})
		return
	}
	c.JSON(http.StatusOK, gin.H{MessageField: SuccessBookUpdated})
}

func (bookHandler *BookHandler) StoreStatsHandler(c *gin.Context) {
	_, ok := utilites.GetUsernameFromQuery(c)
	if !ok {
		return
	}
	bookCount, authorCount, err := bookHandler.bookService.GetStoreStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ErrorField: ErrFailedToGetStats})
		return
	}
	c.JSON(http.StatusOK, gin.H{BookCountField: bookCount, DistinctAuthorsField: authorCount})
}

func (bookHandler *BookHandler) SearchBooksHandler(c *gin.Context) {
	title := c.Query(TitleParam)
	authorName := c.Query(AuthorParam)
	priceRange := c.Query(PriceRangeParam)
	_, ok := utilites.GetUsernameFromQuery(c)
	if !ok {
		return
	}
	if title == EmptyString && authorName == EmptyString && priceRange == EmptyString {
		c.JSON(http.StatusBadRequest, gin.H{ErrorField: ErrParameterRequired})
		return
	}
	if priceRange != EmptyString {
		if utilites.CheckPriceRangeValidity(priceRange) != nil {
			c.JSON(http.StatusBadRequest, gin.H{ErrorField: ErrInvalidPriceRange})
			return
		}
	}
	books, err := bookHandler.bookService.SearchBooks(title, authorName, priceRange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ErrorField: ErrFailedToSearchBooks})
		return
	}
	c.JSON(http.StatusOK, books)
}
