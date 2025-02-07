package models

import "time"

type CrawledPage struct {
	Title          string    `bson:"title"`
	MetaData       string    `bson:"metaData"`
	Url            string    `bson:"url"`
	CrawledAt      time.Time `bson:"crawledAt"`
	TextContent    string    `bson:"textContent"`
	Images         []string  `bson:"images"`
	AssociatedURLs []string  `bson:"associatedURLs"`
}
