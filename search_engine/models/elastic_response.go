package models

import "time"

type ElasticResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index   string   `json:"_index"`
			ID      string   `json:"_id"`
			Score   float64  `json:"_score"`
			Ignored []string `json:"_ignored"`
			Source  struct {
				Title          string    `json:"Title"`
				MetaData       string    `json:"MetaData"`
				URL            string    `json:"Url"`
				CrawledAt      time.Time `json:"CrawledAt"`
				TextContent    string    `json:"TextContent"`
				Images         []string  `json:"Images"`
				AssociatedURLs []string  `json:"AssociatedURLs"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
