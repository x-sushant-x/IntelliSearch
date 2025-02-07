package main

import "github.com/x-sushant-x/IntelliSearch/search_engine/api"

func main() {
	server := api.NewServer("8081")
	server.Start()
}
