package core

import (
	"github.com/x-sushant-x/IntelliSearch/crawler/models"
	"golang.org/x/net/html"
	"log"
	"net/url"
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
			switch node.Data {
			case "header", "nav", "footer", "aside", "form", "script", "noscript", "iframe":
				return
			}
		}

		if node.Type == html.ElementNode {
			if node.Type == html.ElementNode &&
				(node.Data == "p" ||
					node.Data == "h1" ||
					node.Data == "h2" ||
					node.Data == "h3" ||
					node.Data == "h4" ||
					node.Data == "h5" ||
					node.Data == "h6" ||
					node.Data == "div" ||
					node.Data == "span") {

				for _, attr := range node.Attr {
					if attr.Key == "id" || attr.Key == "class" {
						nodeData := attr.Val

						if strings.Contains(nodeData, "menu") ||
							strings.Contains(nodeData, "navigation") ||
							strings.Contains(nodeData, "header") ||
							strings.Contains(nodeData, "footer") {
						}
					}
				}

				crawledContent.TextContent += getTextContent(node) + "\n"
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
						absoluteURL := resolveURL(pageURL, attr.Val)
						crawledContent.AssociatedURLs = append(crawledContent.AssociatedURLs, absoluteURL)
					}
				}
			}
		}

		if node.Type == html.ElementNode && node.Data == "button" {
			return
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			getPageContent(c)
		}
	}

	getPageContent(node)

	return &crawledContent, nil
}

func getTextContent(node *html.Node) string {
	if node.Type == html.TextNode {
		return strings.TrimSpace(node.Data)
	}

	text := ""

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		text += getTextContent(c) + " "
	}

	return strings.TrimSpace(text)
}

func resolveURL(base, relative string) string {
	u, err := url.Parse(relative)
	if err != nil || u.Scheme == "" {
		baseURL, _ := url.Parse(base)
		return baseURL.ResolveReference(u).String()
	}
	return relative
}
