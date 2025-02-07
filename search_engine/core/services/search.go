package services

import (
	database "github.com/x-sushant-x/IntelliSearch/search_engine/core/repository"
	"github.com/x-sushant-x/IntelliSearch/search_engine/models"
)

type SearchService struct {
	db database.DB
}

func NewSearchService(db database.DB) SearchService {
	return SearchService{
		db: db,
	}
}

func (s SearchService) GetSearchResults(query string) (*[]models.SearchResponse, error) {
	return s.db.Search(query)
}
