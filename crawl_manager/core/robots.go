package core

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RobotsChecker struct{}

func (r RobotsChecker) getRobotsFile(link string) ([]string, error) {
	// Complete URL: https://en.wikipedia.org/wiki/Go_(programming_language)

	// Base URL: https://en.wikipedia.org/

	URL, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	host := URL.Host
	scheme := URL.Scheme
	robots := "/robots.txt"

	robotsFilePath := scheme + "://" + host + robots

	resp, err := http.Get(robotsFilePath)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	disallowedLinks := r.parseFile(string(content))

	for _, link := range disallowedLinks {
		fmt.Println("Link: " + link)
	}

	return nil, nil
}

func (r RobotsChecker) parseFile(file string) []string {
	var disallowedLinks []string

	scanner := bufio.NewScanner(strings.NewReader(file))

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 9 {
			s := line[0:9]

			if strings.Contains(s, "Disallow:") {
				split := strings.Split(line, ":")
				link := split[1]
				disallowedLinks = append(disallowedLinks, link)
			}
		}
	}

	return disallowedLinks
}
