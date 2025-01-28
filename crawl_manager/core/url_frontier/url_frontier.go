package urlfrontier

import "github.com/x-sushant-x/IntelliSearch/crawl_manager/core/queue"

type URLFrontier struct {
	messageQueue queue.MessageQueue
}

func NewURLFrontier(queue queue.MessageQueue) URLFrontier {
	return URLFrontier{
		messageQueue: queue,
	}
}
