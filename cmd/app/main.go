package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"grpc-story-service/internal/app/controllers"
	"log"
	"net"
	"net/http"

	"grpc-story-service/internal/app"
	// "google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {

	err := godotenv.Load()

	// Declare an instance of the application
	a := app.New()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	routes := controllers.NewHttpRoutes(a)
	err = http.Serve(lis, routes.Routes())
	if err != nil {
		return
	}

}
