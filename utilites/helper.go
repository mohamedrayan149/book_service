package utilites

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/olivere/elastic/v7"
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
	ErrInvalidPriceRangeFormat = "invalid price range format, expected 'min-max'"
	ErrInvalidPriceRangeValues = "invalid price range values, expected numbers"
)

type BookUpdateData struct {
	ID    string `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
}

func BindBookData(c *gin.Context) (*model.Book, bool) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
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
				default:
					errorMessages[fieldErr.Field()] = ErrFieldRequired
				}
			}
			c.JSON(http.StatusBadRequest, gin.H{ErrorField: errorMessages})
		}
		return nil, false
	}
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

	n1, err1 := strconv.ParseFloat(parts[0], ParseBitSize)
	n2, err2 := strconv.ParseFloat(parts[1], ParseBitSize)

	if err1 != nil || err2 != nil || n1 > n2 {
		return errors.New(ErrInvalidPriceRangeValues)
	}

	return nil
}

func BuildSearchQuery(title, authorName, priceRange string) *elastic.BoolQuery {
	query := elastic.NewBoolQuery()

	if title != EmptyString {
		query.Must(elastic.NewMatchQuery(FieldTitle, title))
	}
	if authorName != EmptyString {
		query.Must(elastic.NewMatchQuery(FieldAuthorName, authorName))
	}
	if priceRange != EmptyString {
		prices := strings.Split(priceRange, PriceRangeDelimiter)
		if len(prices) == 2 { // Ensure price range has both min and max
			from, _ := strconv.ParseFloat(prices[0], ParseBitSize)
			to, _ := strconv.ParseFloat(prices[1], ParseBitSize)
			query.Must(elastic.NewRangeQuery(FieldPrice).Gte(from).Lte(to))
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
