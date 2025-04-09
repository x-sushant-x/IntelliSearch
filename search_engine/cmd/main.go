package main

import (
	"github.com/joho/godotenv"
	"github.com/x-sushant-x/IntelliSearch/search_engine/api"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to lead .env: " + err.Error())
	}
}

func main() {
	server := api.NewServer("8081")
	server.Start()
}
