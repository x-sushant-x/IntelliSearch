package database

import "github.com/x-sushant-x/IntelliSearch/crawl_manager/models"

type DB interface {
	SaveCrawledPage(page *models.CrawledPage)
}
