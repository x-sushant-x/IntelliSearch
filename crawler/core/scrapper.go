package core

import "github.com/gocolly/colly"

func ScrapURL(url string) string {
	html := ""

	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		html = string(r.Body)
	})

	err := c.Visit(url)
	if err != nil {
		return ""
	}

	return html
}
