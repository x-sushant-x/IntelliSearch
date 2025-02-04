package models

import "time"

type CrawledPage struct {
	Title          string
	MetaData       string
	Url            string
	CrawledAt      time.Time
	TextContent    string
	Images         []string
	AssociatedURLs []string
}
