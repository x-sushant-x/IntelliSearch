package elastic

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/x-sushant-x/IntelliSearch/search_engine/models"
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

func SearchDocuments(query string) (*[]models.SearchResponse, error) {
	eQuery := fmt.Sprintf(`{ "query": { "multi_match": { "query": "%s", "fields": ["Title", "MetaData"] } } }`, query)

	res, err := elasticClient.Search(
		elasticClient.Search.WithIndex(ElasticIndex),
		elasticClient.Search.WithBody(strings.NewReader(eQuery)),
		elasticClient.Search.WithPretty(),
	)

	if err != nil {
		log.Printf("Error executing search query: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error response from Elasticsearch: %s", res.String())
		return nil, err
	}

	var elasticResponse models.ElasticResponse

	err = json.NewDecoder(res.Body).Decode(&elasticResponse)
	if err != nil {
		log.Printf("Error executing search query: %v", err)
		return nil, err
	}

	var results []models.SearchResponse

	for _, result := range elasticResponse.Hits.Hits {

		docData := result.Source

		doc := models.SearchResponse{
			Title:           docData.Title,
			MetaDescription: docData.MetaData,
			Url:             docData.URL,
		}

		results = append(results, doc)
	}

	return &results, nil
}
