package elastic

import (
	"github.com/olivere/elastic/v7"
	"log"
)

var Client *elastic.Client

func InitElasticClient() {
	var err error
	Client, err = elastic.NewClient(elastic.SetURL("http://es-search-7.fiverrdev.com:9200"))
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}
}

/*
func elasticSearchInit() (*elasticsearch.Client, error) {
	// Define Elasticsearch configuration
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://es-search-7.fiverrdev.com:9200",
		},
		Username: "foo", // Elasticsearch username
		Password: "bar", // Elasticsearch password
	}
	// Initialize Elasticsearch client
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	// Return the initialized client and no error
	return es, nil
}
*/
