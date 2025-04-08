package elastic

import (
	"bytes"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	elasticClient *elasticsearch.Client
)

const (
	ElasticIndex = "crawled_data"
)

func NewElasticClient() {
	user, pass := getElasticUserAndPassword()

	if elasticClient == nil {
		client, err := elasticsearch.NewClient(elasticsearch.Config{
			Username: user,
			Password: pass,
		})

		if err != nil {
			log.Println("unable to connect to elasticsearch: " + err.Error())
			os.Exit(-1)
		}

		elasticClient = client

		resp, err := client.Ping()
		if err != nil {
			log.Println("error: unable to PING elasticsearch")
			os.Exit(-1)
		}

		if resp.IsError() {
			log.Println("elasticsearch PING returned with error: " + resp.String())
			os.Exit(-1)
		}

		resp, err = client.Indices.Create(ElasticIndex)
		if err != nil {
			log.Println("error: unable to create elasticsearch index")
			os.Exit(-1)
		}

		if resp.IsError() && resp.StatusCode != 400 {
			log.Println("elasticsearch index create returned with error: " + resp.String())
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

func IndexData(data []byte) {
	resp, err := elasticClient.Index(ElasticIndex, bytes.NewReader(data))
	if err != nil {
		log.Println("ERROR: unable to index data: " + err.Error() + " response: " + resp.String())
	}

	if resp.IsError() {
		log.Println("ERROR: Got response as error while indexing data: " + resp.String())
	}

	log.Println("Indexing Response: " + resp.String())
}
