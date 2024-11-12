package connectors

import (
	"github.com/olivere/elastic/v7"
	"library/config"
	"log"
)

func InitElasticClient() *elastic.Client {
	var err error
	ElasticClient, err := elastic.NewClient(
		elastic.SetURL(config.ElasticDevURL),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf(config.ErrorInitClient, err)
	}
	return ElasticClient
}
