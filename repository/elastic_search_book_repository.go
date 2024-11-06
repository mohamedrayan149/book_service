package repository

import (
	"context"
	elasticRemote "github.com/olivere/elastic/v7"
	"library/elastic"
	"library/model"
	"library/utilites"
)

// Constants for Elasticsearch
const (
	IndexBooks         = "books_mohamed"
	FieldTitle         = "title"
	FieldAuthorNameKey = "author_name.keyword"
	AggDistinctAuthors = "distinct_authors"
	EmptyString        = ""
	Zero               = 0
)

type ElasticsearchBookRepository struct {
}

func NewElasticsearchBookRepository() *ElasticsearchBookRepository {
	return &ElasticsearchBookRepository{}
}

func (elasticsearchBookRepository *ElasticsearchBookRepository) GetBookByID(id string) (*model.Book, error) {
	res, err := elastic.Client.Get().Index(IndexBooks).Id(id).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return utilites.ParseBook(res.Source, id)
}

func (elasticsearchBookRepository *ElasticsearchBookRepository) DeleteBookByID(id string) error {
	_, err := elastic.Client.Delete().Index(IndexBooks).Id(id).Do(context.Background())
	return err
}

func (elasticsearchBookRepository *ElasticsearchBookRepository) UpdateBook(id string, title string) error {
	_, err := elastic.Client.Update().
		Index(IndexBooks).Id(id).
		Doc(map[string]interface{}{FieldTitle: title}).
		Do(context.Background())
	return err
}

func (elasticsearchBookRepository *ElasticsearchBookRepository) AddBook(book *model.Book) (string, error) {
	res, err := elastic.Client.Index().Index(IndexBooks).BodyJson(book).Do(context.Background())
	if err != nil {
		return EmptyString, err
	}
	return res.Id, nil
}

func (elasticsearchBookRepository *ElasticsearchBookRepository) GetStoreStats() (int, int, error) {
	numberOfBooks, err := elastic.Client.Count(IndexBooks).Do(context.Background())
	if err != nil {
		return Zero, Zero, err
	}

	agg := elasticRemote.NewCardinalityAggregation().Field(FieldAuthorNameKey)
	searchResult, err := elastic.Client.Search(IndexBooks).
		Aggregation(AggDistinctAuthors, agg).
		Do(context.Background())
	if err != nil {
		return Zero, Zero, err
	}
	aggResult, found := searchResult.Aggregations.Cardinality(AggDistinctAuthors)
	if !found {
		return int(numberOfBooks), Zero, nil
	}

	distinctAuthors := int(*aggResult.Value)
	return int(numberOfBooks), distinctAuthors, nil
}

func (elasticsearchBookRepository *ElasticsearchBookRepository) SearchBooks(title, authorName, priceRange string) ([]*model.Book, error) {
	query := utilites.BuildSearchQuery(title, authorName, priceRange)
	searchResult, err := elastic.Client.Search().Index(IndexBooks).Query(query).Do(context.Background())
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for _, hit := range searchResult.Hits.Hits {
		book, err := utilites.ParseBook(hit.Source, hit.Id)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
