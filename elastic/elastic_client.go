package elastic

import (
	"github.com/olivere/elastic/v7"
	"log"
)

var Client *elastic.Client

func InitElasticClient() {
	var err error
	Client, err = elastic.NewClient(elastic.SetURL("http://es-search-7.fiverrdev.com:9200"), elastic.SetSniff(false))

	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}
}

//http://es-search-7.fiverrdev.com:9200
