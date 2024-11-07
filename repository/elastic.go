package repository

import (
	"context"
	"github.com/olivere/elastic/v7"
	"library/connectors"
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

type BookElasticRepo struct {
	elasticClient *elastic.Client
}

func NewBookElasticRepo() *BookElasticRepo {
	client := connectors.InitElasticClient()
	return &BookElasticRepo{elasticClient: client}
}

func (elasticSearchBook *BookElasticRepo) GetBookByID(id string) (*model.Book, error) {
	res, err := elasticSearchBook.elasticClient.Get().Index(IndexBooks).Id(id).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return utilites.ParseBook(res.Source, id)
}

func (elasticSearchBook *BookElasticRepo) DeleteBookByID(id string) error {
	_, err := elasticSearchBook.elasticClient.Delete().Index(IndexBooks).Id(id).Do(context.Background())
	return err
}

func (elasticSearchBook *BookElasticRepo) UpdateBook(id string, title string) error {
	_, err := elasticSearchBook.elasticClient.Update().
		Index(IndexBooks).Id(id).
		Doc(map[string]interface{}{FieldTitle: title}).
		Do(context.Background())
	return err
}

func (elasticSearchBook *BookElasticRepo) AddBook(book *model.Book) (string, error) {
	res, err := elasticSearchBook.elasticClient.Index().Index(IndexBooks).BodyJson(book).Do(context.Background())
	if err != nil {
		return EmptyString, err
	}
	return res.Id, nil
}

func (elasticSearchBook *BookElasticRepo) GetStoreStats() (int, int, error) {
	numberOfBooks, err := elasticSearchBook.elasticClient.Count(IndexBooks).Do(context.Background())
	if err != nil {
		return Zero, Zero, err
	}

	agg := elastic.NewCardinalityAggregation().Field(FieldAuthorNameKey)
	searchResult, err := elasticSearchBook.elasticClient.Search(IndexBooks).
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

func (elasticSearchBook *BookElasticRepo) SearchBooks(title, authorName, priceRange string) ([]*model.Book, error) {
	query := utilites.BuildSearchQuery(title, authorName, priceRange)
	searchResult, err := elasticSearchBook.elasticClient.Search().Index(IndexBooks).Query(query).Do(context.Background())
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
