package urlfrontier

import (
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/queue"
	"log"
)

const (
	crawlURLsKafkaTopic = "crawl_urls"
)

type URLFrontier struct {
	messageQueue  queue.MessageQueue
	onGoingCrawls map[string]bool
}

func NewURLFrontier(queue queue.MessageQueue) URLFrontier {
	return URLFrontier{
		messageQueue: queue,
	}
}

func (frontier URLFrontier) SendURLToQueueForCrawling(urls []string) {
	for _, url := range urls {
		err := frontier.messageQueue.Send(crawlURLsKafkaTopic, "", url)
		if err != nil {
			log.Default().Println("Unable to send crawling URL to queue:", err)
		}
	}
}
