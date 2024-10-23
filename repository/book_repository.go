package repository

import (
	"context"
	"encoding/json"
	"library/elastic"
	"library/model"
)

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (br *BookRepository) GetBookByID(id string) (*model.Book, error) {
	res, err := elastic.Client.Get().Index("books_mohamed").Id(id).Do(context.Background())
	if err != nil {
		return nil, err
	}
	var book model.Book
	err = json.Unmarshal(res.Source, &book)
	return &book, err
}

func (br *BookRepository) DeleteBookByID(id string) error {
	_, err := elastic.Client.Delete().Index("books_mohamed").Id(id).Do(context.Background())
	return err
}
