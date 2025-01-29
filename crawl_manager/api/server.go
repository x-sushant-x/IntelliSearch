package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/queue"
	urlfrontier "github.com/x-sushant-x/IntelliSearch/crawl_manager/core/url_frontier"
	"log"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{port: port}
}

func (s *Server) Start(messageQueue queue.MessageQueue) {
	app := fiber.New()

	frontier := urlfrontier.NewURLFrontier(messageQueue)
	crawlController := NewCrawlController(frontier)

	app.Post("/crawl", crawlController.HandleNewCrawlRequest)

	err := app.Listen(":" + s.port)
	if err != nil {
		log.Fatal("unable to start server: " + err.Error())
	}
}
