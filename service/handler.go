package service

import (
	"github.com/gin-gonic/gin"
	"library/repository"
	"library/utilites"
	"net/http"
)

// Constants for parameter names
const (
	IDParam         = "id"
	TitleParam      = "title"
	AuthorParam     = "author_name"
	PriceRangeParam = "price_range"
	UsernameParam   = "username"
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
	ErrorUsernameRequired  = "username is required"
)

// Constants for success messages
const (
	SuccessBookDeleted = "Book deleted successfully"
	SuccessBookUpdated = "Book updated successfully"
)

// Constants for response fields
const (
	ActionsField         = "actions"
	BookCountField       = "book_count"
	DistinctAuthorsField = "distinct_authors"
	MessageField         = "message"
	ErrorField           = "error"
)

type Handler struct {
	bookRepository     repository.BookRepository
	activityRepository repository.ActivityRepository
}

func NewHandler(bookRepository repository.BookRepository, activityRepository repository.ActivityRepository) *Handler {
	return &Handler{bookRepository: bookRepository, activityRepository: activityRepository}
}

func (h *Handler) GetBook(c *gin.Context) {
	id, ok := utilites.GetIDFromQuery(c)
	if !ok {
		return // The error response is already handled by GetIDFromQuery
	}
	book, err := h.bookRepository.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{ErrorField: ErrFailedToGetBook})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *Handler) DeleteBook(c *gin.Context) {
	id, ok := utilites.GetIDFromQuery(c)
	if !ok {
		return // The error response is already handled by GetIDFromQuery
	}
	err := h.bookRepository.DeleteBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{ErrorField: ErrFailedToDeleteBook})
		return
	}
	c.JSON(http.StatusOK, gin.H{MessageField: SuccessBookDeleted})
}

func (h *Handler) AddBook(c *gin.Context) {
	book, ok := utilites.BindBookData(c)
	if !ok {
		return
	}
	id, err := h.bookRepository.AddBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ErrorField: ErrFailedToAddBook})
		return
	}
	c.JSON(http.StatusCreated, gin.H{IDParam: id})
}

func (h *Handler) UpdateBook(c *gin.Context) {
	data, ok := utilites.BindBookUpdateData(c)
	if !ok {
		return
	}
	err := h.bookRepository.UpdateBook(data.ID, data.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ErrorField: ErrFailedToUpdateBook})
		return
	}
	c.JSON(http.StatusOK, gin.H{MessageField: SuccessBookUpdated})
}

func (h *Handler) StoreStats(c *gin.Context) {
	bookCount, authorCount, err := h.bookRepository.GetStoreStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ErrorField: ErrFailedToGetStats})
		return
	}
	c.JSON(http.StatusOK, gin.H{BookCountField: bookCount, DistinctAuthorsField: authorCount})
}

func (h *Handler) SearchBooks(c *gin.Context) {
	title := c.Query(TitleParam)
	authorName := c.Query(AuthorParam)
	priceRange := c.Query(PriceRangeParam)

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
	books, err := h.bookRepository.SearchBooks(title, authorName, priceRange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ErrorField: ErrFailedToSearchBooks})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *Handler) Activity(c *gin.Context) {
	username := c.Query(UsernameParam)
	if username == EmptyString {
		c.JSON(http.StatusBadRequest, gin.H{ErrorField: ErrorUsernameRequired})
		return
	}

	actions, err := h.activityRepository.GetLastUserActions(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ErrorField: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{ActionsField: actions})
}
