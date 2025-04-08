package core

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	ElasticClient *elasticsearch.Client
)

func NewElasticClient() {
	user, pass := getElasticUserAndPassword()

	if ElasticClient == nil {
		client, err := elasticsearch.NewClient(elasticsearch.Config{
			Username: user,
			Password: pass,
		})

		if err != nil {
			log.Println("unable to connect to elasticsearch: " + err.Error())
			os.Exit(-1)
		}

		ElasticClient = client

		resp, err := client.Ping()
		if err != nil {
			log.Println("error: unable to PING elasticsearch")
			os.Exit(-1)
		}

		if resp.IsError() {
			log.Println("elasticsearch PING returned with error: " + resp.String())
			os.Exit(-1)
		}

		log.Println("Elasticsearch initialized successfully")
	} else {
		log.Println("ElasticSearch already initialized")
	}
}

func getElasticUserAndPassword() (string, string) {
	user := os.Getenv("ELASTIC_USER")
	pass := os.Getenv("ELASTIC_PASSWORD")

	if len(user) == 0 {
		log.Println("error: ELASTIC_USER not available in env variables")
		os.Exit(-1)
	}

	if len(pass) == 0 {
		log.Println("error: ELASTIC_PASSWORD not available in env variables")
		os.Exit(-1)
	}

	return user, pass
}
