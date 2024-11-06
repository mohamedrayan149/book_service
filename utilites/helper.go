package utilites

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	elasticRemote "github.com/olivere/elastic/v7"
	"library/model"
	"net/http"
	"strconv"
	"strings"
)

const (
	FieldID          = "id"
	FieldTitle       = "title"
	FieldAuthorName  = "author_name"
	FieldPrice       = "price"
	FieldUserName    = "username"
	FieldPublishDate = "publish_date"
	ErrorField       = "error"
)
const (
	PriceRangeDelimiter = "-"
	ParseBitSize        = 64
	EmptyString         = ""
	ID                  = "ID"
	Title               = "Title"
	AuthorName          = "AuthorName"
	UserName            = "UserName"
	PublishDate         = "PublishDate"
	Price               = "Price"
)

const (
	ErrIDRequired              = "ID is a required field."
	ErrTitleRequired           = "Title is a required field."
	ErrAuthorNameRequired      = "Author name is a required field."
	ErrPriceRequired           = "Price is a required field and must be a number."
	ErrFieldRequired           = "This field is required."
	ErrPublishDateRequired     = "Publish date is a required field."
	ErrUserNameRequired        = "Username is a required field."
	ErrInvalidPriceRangeFormat = "invalid price range format, expected 'min-max'"
	ErrInvalidPriceRangeValues = "invalid price range values, expected numbers"
)

type BookUpdateData struct {
	ID       string `json:"id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Username string `json:"username" binding:"required"`
}
type BookWithUserName struct {
	ID             string  `json:"id"`
	Title          string  `json:"title" binding:"required"`
	AuthorName     string  `json:"author_name" binding:"required"`
	Price          float64 `json:"price" binding:"required"`
	EbookAvailable bool    `json:"ebook_available"`
	PublishDate    string  `json:"publish_date" binding:"required"`
	UserName       string  `json:"username" binding:"required"`
}

func BindBookData(c *gin.Context) (*model.Book, bool) {
	var data BookWithUserName
	if err := c.ShouldBindJSON(&data); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			errorMessages := make(map[string]string)
			for _, fieldErr := range validationErrors {
				switch fieldErr.Field() {
				case Title:
					errorMessages[FieldTitle] = ErrTitleRequired
				case AuthorName:
					errorMessages[FieldAuthorName] = ErrAuthorNameRequired
				case Price:
					errorMessages[FieldPrice] = ErrPriceRequired
				case PublishDate:
					errorMessages[FieldPublishDate] = ErrPublishDateRequired
				case UserName:
					errorMessages[FieldUserName] = ErrUserNameRequired
				default:
					errorMessages[fieldErr.Field()] = ErrFieldRequired
				}
			}
			c.JSON(http.StatusBadRequest, gin.H{ErrorField: errorMessages})
		}
		return nil, false
	}
	var book model.Book
	book.ID = data.ID
	book.Title = data.Title
	book.AuthorName = data.AuthorName
	book.Price = data.Price
	book.EbookAvailable = data.EbookAvailable
	book.PublishDate = data.PublishDate

	return &book, true
}

func BindBookUpdateData(c *gin.Context) (*BookUpdateData, bool) {
	var data BookUpdateData
	if err := c.ShouldBindJSON(&data); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			errorMessages := make(map[string]string)
			for _, fieldErr := range validationErrors {
				switch fieldErr.Field() {
				case ID:
					errorMessages[FieldID] = ErrIDRequired
				case Title:
					errorMessages[FieldTitle] = ErrTitleRequired
				case UserName:
					errorMessages[FieldUserName] = ErrUserNameRequired
				default:
					errorMessages[fieldErr.Field()] = ErrFieldRequired
				}
			}
			c.JSON(http.StatusBadRequest, gin.H{ErrorField: errorMessages})
		}
		return nil, false
	}
	return &data, true
}

func CheckPriceRangeValidity(priceRange string) error {
	parts := strings.Split(priceRange, PriceRangeDelimiter)
	if len(parts) != 2 {
		return errors.New(ErrInvalidPriceRangeFormat)
	}

	_, err1 := strconv.ParseFloat(parts[0], ParseBitSize)
	_, err2 := strconv.ParseFloat(parts[1], ParseBitSize)

	if err1 != nil || err2 != nil {
		return errors.New(ErrInvalidPriceRangeValues)
	}

	return nil
}

func BuildSearchQuery(title, authorName, priceRange string) *elasticRemote.BoolQuery {
	query := elasticRemote.NewBoolQuery()

	if title != EmptyString {
		query.Must(elasticRemote.NewMatchQuery(FieldTitle, title))
	}
	if authorName != EmptyString {
		query.Must(elasticRemote.NewMatchQuery(FieldAuthorName, authorName))
	}
	if priceRange != EmptyString {
		prices := strings.Split(priceRange, PriceRangeDelimiter)
		if len(prices) == 2 { // Ensure price range has both min and max
			from, _ := strconv.ParseFloat(prices[0], ParseBitSize)
			to, _ := strconv.ParseFloat(prices[1], ParseBitSize)
			query.Must(elasticRemote.NewRangeQuery(FieldPrice).Gte(from).Lte(to))
		}
	}
	return query
}

func ParseBook(source []byte, id string) (*model.Book, error) {
	var book model.Book
	err := json.Unmarshal(source, &book)
	if err != nil {
		return nil, err
	}
	book.ID = id
	return &book, nil
}

func GetIDFromQuery(c *gin.Context) (string, bool) {
	id := c.Query(FieldID)
	if id == EmptyString {
		c.JSON(http.StatusBadRequest, gin.H{ErrorField: ErrIDRequired})
		return EmptyString, false
	}
	return id, true
}

func GetUsernameFromQuery(c *gin.Context) (string, bool) {
	username := c.Query(FieldUserName)
	if username == EmptyString {
		c.JSON(http.StatusBadRequest, gin.H{ErrorField: ErrUserNameRequired})
		return EmptyString, false
	}
	return username, true
}
