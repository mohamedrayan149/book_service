package repository

import (
	"context"
	"encoding/json"
	elasticRemote "github.com/olivere/elastic/v7"
	"library/elastic"
	"library/model"
	"strconv"
	"strings"
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
	book.ID = id
	return &book, err
}
func (br *BookRepository) DeleteBookByID(id string) error {
	_, err := elastic.Client.Delete().Index("books_mohamed").Id(id).Do(context.Background())
	return err
}
func (br *BookRepository) UpdateBook(id string, title string) error {
	_, err := elastic.Client.Update().
		Index("books_mohamed").Id(id).
		Doc(map[string]interface{}{"title": title}).
		Do(context.Background())
	return err
}
func (br *BookRepository) AddBook(book *model.Book) (string, error) {
	res, err := elastic.Client.Index().Index("books_mohamed").BodyJson(book).Do(context.Background())
	if err != nil {
		return "", err
	}
	return res.Id, nil
}
func (br *BookRepository) GetStoreStats() (int, int, error) {
	numberOfBooks, err := elastic.Client.Count("books_mohamed").Do(context.Background())
	if err != nil {
		return 0, 0, err
	}

	agg := elasticRemote.NewCardinalityAggregation().Field("author_name.keyword")
	searchResult, err := elastic.Client.Search("books_mohamed").
		Aggregation("distinct_authors", agg).
		Do(context.Background())
	if err != nil {
		return 0, 0, err
	}
	aggResult, found := searchResult.Aggregations.Cardinality("distinct_authors")
	if !found {
		return int(numberOfBooks), 0, nil
	}

	distinctAuthors := int(*aggResult.Value)
	return int(numberOfBooks), distinctAuthors, nil
}
func (br *BookRepository) SearchBooks(title, authorName, priceRange string) ([]*model.Book, error) {
	query := elasticRemote.NewBoolQuery()
	if title != "" {
		query.Must(elasticRemote.NewMatchQuery("title", title))
	}
	if authorName != "" {
		query.Must(elasticRemote.NewMatchQuery("author_name", authorName))
	}
	if priceRange != "" {
		prices := strings.Split(priceRange, "-")
		if len(prices) == 2 {
			from, err1 := strconv.ParseFloat(prices[0], 64)
			to, err2 := strconv.ParseFloat(prices[1], 64)
			if err1 == nil && err2 == nil {
				query.Must(elasticRemote.NewRangeQuery("price").Gte(from).Lte(to))
			}
		}
	}
	searchResult, err := elastic.Client.Search().Index("books_mohamed").Query(query).Do(context.Background())
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for _, hit := range searchResult.Hits.Hits {
		var book model.Book
		book.ID = hit.Id
		err := json.Unmarshal(hit.Source, &book)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}
