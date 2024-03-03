package main

import (
	"github.com/joho/godotenv"
	"log"

	"grpc-story-service/internal/app"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.StartServer()
}
