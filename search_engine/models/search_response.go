package models

type SearchResponse struct {
	Title           string `json:"title"`
	MetaDescription string `json:"metaData"`
	Url             string `json:"url"`
}
