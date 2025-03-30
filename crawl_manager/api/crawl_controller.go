package api

import (
	"github.com/gofiber/fiber/v2"
	urlfrontier "github.com/x-sushant-x/IntelliSearch/crawl_manager/core"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/models"
)

type CrawlController struct {
	urlFrontier urlfrontier.URLFrontier
}

func NewCrawlController(urlFrontier urlfrontier.URLFrontier) *CrawlController {
	return &CrawlController{urlFrontier}
}

func (c *CrawlController) HandleNewCrawlRequest(ctx *fiber.Ctx) error {
	var reqBody models.NewCrawlRequest

	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	c.urlFrontier.SendURLToQueueForCrawling(reqBody.URLs)

	return ctx.Status(fiber.StatusOK).JSON(map[string]string{
		"message": "Crawl Started Successfully",
	})
}
