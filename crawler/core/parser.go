package core

import (
	"golang.org/x/net/html"
	"log"
	"strings"
)

func ExtractTitleAndMetaData(htmlContent string) (string, string, error) {
	node, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Println("error while parsing html: " + err.Error())
		return "", "", err
	}

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
