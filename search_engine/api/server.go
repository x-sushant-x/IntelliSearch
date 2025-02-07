package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	database "github.com/x-sushant-x/IntelliSearch/search_engine/core/repository"
	"github.com/x-sushant-x/IntelliSearch/search_engine/core/services"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{port: port}
}

func (s *Server) Start() {
	app := fiber.New()

	db := database.NewMongoDBConnection()

	// db.CreateIndexes()

	svc := services.NewSearchService(db)
	searchHandler := NewSearchHandler(svc)

	app.Get("/search", searchHandler.HandleSearch)

	err := app.Listen(":" + s.port)
	if err != nil {
		log.Fatal("unable to start server: " + err.Error())
	}
}
