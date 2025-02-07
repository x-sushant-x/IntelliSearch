package database

import "github.com/x-sushant-x/IntelliSearch/search_engine/models"

type DB interface {
	Search(query string) (*[]models.SearchResponse, error)
	CreateIndexes()
}
