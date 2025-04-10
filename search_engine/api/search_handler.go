package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/x-sushant-x/IntelliSearch/search_engine/core/services"
	"github.com/x-sushant-x/IntelliSearch/search_engine/elastic"
)

type SearchHandler struct {
	searchService services.SearchService
}

func NewSearchHandler(searchService services.SearchService) SearchHandler {
	return SearchHandler{
		searchService: searchService,
	}
}

func (h SearchHandler) HandleSearch(ctx *fiber.Ctx) error {
	query := ctx.Query("query")
	repoType := ctx.Query("repo_type")

	if repoType == "mongo" {
		resp, err := h.searchService.GetSearchResults(query)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{
				"error": err.Error(),
			})
		}

		return ctx.Status(http.StatusOK).JSON(resp)
	} else if repoType == "elastic" {
		resp, err := elastic.SearchDocuments(query)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{
				"error": err.Error(),
			})
		}

		return ctx.Status(http.StatusOK).JSON(resp)
	}

	return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
		"error": "Please specify repo_type. Expected values can be mongo | elastic",
	})
}
