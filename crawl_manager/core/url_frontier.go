package core

import (
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/cache"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/queue"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/utils"
	"log"
	"strings"
)

const (
	crawlURLsKafkaTopic = "crawl_urls"
)

type URLFrontier struct {
	messageQueue     queue.MessageQueue
	bloomCache       cache.Cache
	robots           map[string]bool // This will tell if a robots file is already fetched and stored in bloom filter. TODO - Change with better apprach later
	robotsDownloader RobotsDownloader
}

func NewURLFrontier(queue queue.MessageQueue, cache cache.Cache) URLFrontier {
	return URLFrontier{
		messageQueue:     queue,
		bloomCache:       cache,
		robots:           make(map[string]bool),
		robotsDownloader: NewRobotsDownloader(),
	}
}

func (f URLFrontier) SendURLToQueueForCrawling(urls []string) {
	for _, url := range urls {
		_, found := f.robots[url]

		if !found {
			hostName, err := utils.GetURLHostName(url)
			if err != nil {
				log.Println("error: Invalid URL")
				continue
			}

			log.Printf("Downloading robots.txt for host name: %s\n", hostName)
			disallowedLinks, err := f.robotsDownloader.GetDisallowedLinks(url)
			if err != nil {
				continue
			}

			for _, link := range disallowedLinks {
				link = strings.TrimSpace(link)
				url = strings.TrimSpace(url)
				f.bloomCache.InsertBloomFilter(url + link)

				f.robots[hostName] = true
			}
		}
	}

	for _, url := range urls {
		hostName, err := utils.GetURLHostName(url)
		if err != nil {
			log.Println("error: Invalid URL")
			continue
		}

		_, found := f.robots[hostName]

		if found {
			log.Printf("Cached robots.txt found for host name: %s\n", hostName)
			url = strings.TrimSpace(url)

			disallowed := f.bloomCache.CheckBloom(url)

			if disallowed {
				log.Println("Invalid Crawl: This page is not allowed to be crawled. Check robots.txt.")
				continue
			}

			log.Printf("Crawl Started: Sent url %s for crawling.\n", url)
			err := f.messageQueue.Send(crawlURLsKafkaTopic, "", url)
			if err != nil {
				log.Default().Println("Unable to send crawling URL to queue:", err)
			}
		}

	}
}
