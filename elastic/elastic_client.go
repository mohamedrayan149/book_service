package elastic

import (
	"github.com/olivere/elastic/v7"
	"log"
)

const (
	ElasticDevURL   = "http://es-search-7.fiverrdev.com:9200"
	ErrorInitClient = "Error creating Elasticsearch client: %s"
)

var Client *elastic.Client

func InitElasticClient() {
	var err error
	Client, err = elastic.NewClient(
		elastic.SetURL(ElasticDevURL),
		elastic.SetSniff(false),
	)

	if err != nil {
		log.Fatalf(ErrorInitClient, err)
	}
}
