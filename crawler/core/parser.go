package core

import (
	"github.com/x-sushant-x/IntelliSearch/crawler/models"
	"golang.org/x/net/html"
	"log"
	"strings"
	"time"
)

func extractTitleAndMetaData(node *html.Node) (string, string, error) {
	var title, metaDescription string

	var findMetaDescriptionAndTitle func(n *html.Node)
	findMetaDescriptionAndTitle = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = n.FirstChild.Data
		}

		if n.Type == html.ElementNode && n.Data == "meta" {
			var name, content string

			for _, a := range n.Attr {
				if a.Key == "name" && a.Val == "description" {
					name = a.Val
				}

				if a.Key == "content" {
					content = a.Val
				}
			}
			if name == "description" {
				metaDescription = content
				return
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findMetaDescriptionAndTitle(c)
		}
	}

	findMetaDescriptionAndTitle(node)

	return title, metaDescription, nil
}

func ExtractContent(htmlContent, pageURL string) (*models.CrawledPage, error) {
	crawledContent := models.CrawledPage{}

	crawledContent.CrawledAt = time.Now()
	crawledContent.Url = pageURL

	node, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Println("error while parsing html: " + err.Error())
		return nil, err
	}

	title, metaData, err := extractTitleAndMetaData(node)
	if err != nil {
		log.Println("error while extracting title and metadata: " + err.Error())
		return nil, err
	}

	crawledContent.Title = title
	crawledContent.MetaData = metaData

	var getPageContent func(node *html.Node)

	getPageContent = func(node *html.Node) {
		if node.Type == html.ElementNode {
			if node.Data == "p" && node.FirstChild != nil {
				crawledContent.TextContent += node.FirstChild.Data + "\n"
			}

			if node.Type == html.ElementNode &&
				(node.Data == "h1" ||
					node.Data == "h2" ||
					node.Data == "h3" ||
					node.Data == "h4" ||
					node.Data == "h5" ||
					node.Data == "h6") {
				crawledContent.TextContent += node.FirstChild.Data + "\n"
			}

			if node.Data == "img" {
				for _, attr := range node.Attr {
					if attr.Key == "src" {
						imageURL := attr.Val
						crawledContent.Images = append(crawledContent.Images, imageURL)

					}
				}
			}

			if node.Data == "a" {
				for _, attr := range node.Attr {
					if attr.Key == "href" {
						URL := attr.Val
						crawledContent.AssociatedURLs = append(crawledContent.AssociatedURLs, URL)
					}
				}
			}
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			getPageContent(c)
		}
	}

	getPageContent(node)

	return &crawledContent, nil
}
