package utilities

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/olivere/elastic/v7"
	"library/config"
	"library/model"
	"regexp"
	"strconv"
	"strings"
)

func CheckPriceRangeValidity(priceRange string) error {
	parts := strings.Split(priceRange, config.PriceRangeDelimiter)
	if len(parts) != 2 {
		return errors.New(config.ErrInvalidPriceRangeFormat)
	}

	minimum, err1 := strconv.ParseFloat(parts[0], config.ParseBitSize)
	maximum, err2 := strconv.ParseFloat(parts[1], config.ParseBitSize)

	if err1 != nil || err2 != nil {
		return errors.New(config.ErrInvalidPriceRangeValues)
	}

	if minimum > maximum {
		return errors.New(config.ErrInvalidPriceRangeFormat)
	}

	return nil
}

func BuildSearchQuery(title, authorName, priceRange string) *elastic.BoolQuery {
	query := elastic.NewBoolQuery()

	if title != config.EmptyString {
		query.Must(elastic.NewMatchQuery(config.FieldTitle, title))
	}
	if authorName != config.EmptyString {
		query.Must(elastic.NewMatchQuery(config.FieldAuthorName, authorName))
	}
	if priceRange != config.EmptyString {
		prices := strings.Split(priceRange, config.PriceRangeDelimiter)
		from, _ := strconv.ParseFloat(prices[0], config.ParseBitSize)
		to, _ := strconv.ParseFloat(prices[1], config.ParseBitSize)
		query.Must(elastic.NewRangeQuery(config.FieldPrice).Gte(from).Lte(to))
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

func ParseElasticsearchErrorCode(err error) config.HttpError {
	// Define a regex to capture the status code from the error string
	re := regexp.MustCompile(`Error (\d{3})`)
	// Find the match
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) < 2 {
		return config.HttpError{Code: 0}
	}
	// Convert the code to an integer
	code, convErr := strconv.Atoi(matches[1])
	if convErr != nil {
		return config.HttpError{Code: 0}
	}

	return config.HttpError{Code: code}
}

func GetUpdateBookValidationErrors(err error) map[string]string {
	var validationErrors validator.ValidationErrors
	errorMessages := make(map[string]string)

	if errors.As(err, &validationErrors) {
		for _, fieldErr := range validationErrors {
			switch fieldErr.Field() {
			case config.FieldID:
				errorMessages[config.FieldID] = config.ErrIDRequired
			case config.FieldTitle:
				errorMessages[config.FieldTitle] = config.ErrTitleRequired
			default:
				errorMessages[fieldErr.Field()] = config.ErrFieldRequired
			}
		}
	}

	return errorMessages
}

func GetAddBookValidationErrors(err error) map[string]string {
	var validationErrors validator.ValidationErrors
	errorMessages := make(map[string]string)

	if errors.As(err, &validationErrors) {
		for _, fieldErr := range validationErrors {
			switch fieldErr.Field() {
			case config.FieldTitle:
				errorMessages[config.FieldTitle] = config.ErrTitleRequired
			case config.FieldAuthorName:
				errorMessages[config.FieldAuthorName] = config.ErrAuthorNameRequired
			case config.FieldPrice:
				errorMessages[config.FieldPrice] = config.ErrPriceRequired
			case config.FieldPublishDate:
				errorMessages[config.FieldPublishDate] = config.ErrPublishDateRequired
			default:
				errorMessages[fieldErr.Field()] = config.ErrFieldRequired
			}
		}
	}
	return errorMessages
}
