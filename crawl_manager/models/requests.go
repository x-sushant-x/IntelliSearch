package models

type NewCrawlRequest struct {
	URLs []string `json:"urls"`
}
