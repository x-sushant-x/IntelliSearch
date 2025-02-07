package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/x-sushant-x/IntelliSearch/search_engine/core/services"
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

	resp, err := h.searchService.GetSearchResults(query)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(resp)
}
