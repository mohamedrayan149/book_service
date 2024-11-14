package datastore

import (
	"context"
	"github.com/olivere/elastic/v7"
	"library/config"
	"library/connectors"
	"library/model"
	"library/utilities"
)

type BookStoreElastic struct {
	elasticClient *elastic.Client
}

func NewBookStoreElastic() *BookStoreElastic {
	client := connectors.InitElasticClient()
	return &BookStoreElastic{elasticClient: client}
}

func (bs *BookStoreElastic) GetBook(id string) (*model.Book, error) {
	res, err := bs.elasticClient.Get().Index(config.IndexBooks).Id(id).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return utilities.ParseBook(res.Source, id)
}

func (bs *BookStoreElastic) DeleteBook(id string) error {
	_, err := bs.elasticClient.Delete().Index(config.IndexBooks).Id(id).Do(context.Background())
	return err
}

func (bs *BookStoreElastic) UpdateBook(id string, title string) error {
	_, err := bs.elasticClient.Update().
		Index(config.IndexBooks).Id(id).
		Doc(map[string]interface{}{config.FieldTitle: title}).
		Do(context.Background())
	return err
}

func (bs *BookStoreElastic) AddBook(book *config.BookAddData) (string, error) {
	res, err := bs.elasticClient.Index().Index(config.IndexBooks).BodyJson(book).Do(context.Background())
	if err != nil {
		return config.EmptyString, err
	}
	return res.Id, nil
}

func (bs *BookStoreElastic) GetStoreStats() (int, int, error) {
	numberOfBooks, err := bs.elasticClient.Count(config.IndexBooks).Do(context.Background())
	if err != nil {
		return config.Zero, config.Zero, err
	}

	agg := elastic.NewCardinalityAggregation().Field(config.FieldAuthorNameKey)
	searchResult, err := bs.elasticClient.Search(config.IndexBooks).
		Aggregation(config.AggDistinctAuthors, agg).
		Do(context.Background())
	if err != nil {
		return config.Zero, config.Zero, err
	}
	aggResult, found := searchResult.Aggregations.Cardinality(config.AggDistinctAuthors)
	if !found {
		return int(numberOfBooks), config.Zero, nil
	}

	distinctAuthors := int(*aggResult.Value)
	return int(numberOfBooks), distinctAuthors, nil
}

func (bs *BookStoreElastic) SearchBooks(title, authorName, priceRange string) ([]*model.Book, error) {
	query := utilities.BuildSearchQuery(title, authorName, priceRange)
	searchResult, err := bs.elasticClient.Search().Index(config.IndexBooks).Query(query).Do(context.Background())
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for _, hit := range searchResult.Hits.Hits {
		book, err := utilities.ParseBook(hit.Source, hit.Id)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
