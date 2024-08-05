package main

import (
	"log"

	"github.com/esrichardd/tech-interview-scrappers/internal/adapters/api"
)

func main() {
	api.StartServer()
	log.Println("Server started on port 3003")
}
