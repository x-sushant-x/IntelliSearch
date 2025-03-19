package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New())

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
