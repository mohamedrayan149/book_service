package service

import (
	"github.com/gin-gonic/gin"
	"library/config"
	"library/datastore"
	"library/utilities"
	"net/http"
)

type Handler struct {
	bookStore    datastore.BookStore
	userActivity datastore.UserActivity
}

func NewHandler(bookStore datastore.BookStore, userActivity datastore.UserActivity) *Handler {
	return &Handler{bookStore: bookStore, userActivity: userActivity}
}

func (h *Handler) GetBook(c *gin.Context) {
	id := c.Query(config.FieldID)
	if id == config.EmptyString {
		c.JSON(http.StatusBadRequest, gin.H{config.ErrorField: config.ErrIDRequired})
		return
	}
	book, err := h.bookStore.GetBook(id)
	if err != nil {
		httpError := utilities.ParseElasticsearchErrorCode(err)
		c.JSON(httpError.Code, gin.H{config.ErrorField: config.ElasticsearchErrorMap[httpError.Code]})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *Handler) DeleteBook(c *gin.Context) {
	id := c.Query(config.FieldID)
	if id == config.EmptyString {
		c.JSON(http.StatusBadRequest, gin.H{config.ErrorField: config.ErrIDRequired})
		return
	}
	err := h.bookStore.DeleteBook(id)
	if err != nil {
		httpError := utilities.ParseElasticsearchErrorCode(err)
		c.JSON(httpError.Code, gin.H{config.ErrorField: config.ElasticsearchErrorMap[httpError.Code]})
		return
	}
	c.JSON(http.StatusOK, gin.H{config.MessageField: config.SuccessBookDeleted})
}

func (h *Handler) AddBook(c *gin.Context) {
	var bookToAdd config.BookAddData
	if err := c.ShouldBindJSON(&bookToAdd); err != nil {
		errorMessages := utilities.GetAddBookValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{config.ErrorField: errorMessages})
		return
	}
	id, err := h.bookStore.AddBook(&bookToAdd)
	if err != nil {
		httpError := utilities.ParseElasticsearchErrorCode(err)
		c.JSON(httpError.Code, gin.H{config.ErrorField: config.ElasticsearchErrorMap[httpError.Code]})
		return
	}
	c.JSON(http.StatusCreated, gin.H{config.IDParam: id})
}

func (h *Handler) UpdateBook(c *gin.Context) {
	var bookToUpdate config.BookUpdateData
	if err := c.ShouldBindJSON(&bookToUpdate); err != nil {
		errorMessages := utilities.GetUpdateBookValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{config.ErrorField: errorMessages})
		return
	}
	err := h.bookStore.UpdateBook(bookToUpdate.ID, bookToUpdate.Title)
	if err != nil {
		httpError := utilities.ParseElasticsearchErrorCode(err)
		c.JSON(httpError.Code, gin.H{config.ErrorField: config.ElasticsearchErrorMap[httpError.Code]})
		return
	}
	c.JSON(http.StatusOK, gin.H{config.MessageField: config.SuccessBookUpdated})
}

func (h *Handler) StoreStats(c *gin.Context) {
	bookCount, authorCount, err := h.bookStore.GetStoreStats()
	if err != nil {
		httpError := utilities.ParseElasticsearchErrorCode(err)
		c.JSON(httpError.Code, gin.H{config.ErrorField: config.ElasticsearchErrorMap[httpError.Code]})
		return
	}
	c.JSON(http.StatusOK, gin.H{config.BookCountField: bookCount, config.DistinctAuthorsField: authorCount})
}

func (h *Handler) SearchBooks(c *gin.Context) {
	title := c.Query(config.TitleParam)
	authorName := c.Query(config.AuthorParam)
	priceRange := c.Query(config.PriceRangeParam)

	if title == config.EmptyString && authorName == config.EmptyString && priceRange == config.EmptyString {
		c.JSON(http.StatusBadRequest, gin.H{config.ErrorField: config.ErrParameterRequired})
		return
	}
	if priceRange != config.EmptyString {
		if err := utilities.CheckPriceRangeValidity(priceRange); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{config.ErrorField: err.Error()})
			return
		}
	}
	books, err := h.bookStore.SearchBooks(title, authorName, priceRange)
	if err != nil {
		httpError := utilities.ParseElasticsearchErrorCode(err)
		c.JSON(httpError.Code, gin.H{config.ErrorField: config.ElasticsearchErrorMap[httpError.Code]})
		return
	}
	if len(books) == 0 {
		c.JSON(http.StatusNotFound, config.NoResultFound)
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *Handler) Activity(c *gin.Context) {
	username := c.Query(config.UsernameParam)
	if username == config.EmptyString {
		c.JSON(http.StatusBadRequest, gin.H{config.ErrorField: config.ErrorUsernameRequired})
		return
	}
	actions, err := h.userActivity.GetLastUserActions(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{config.ErrorField: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{config.ActionsField: actions})
}
