package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"grpc-story-service/internal/app"
	// "google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {

	// Declare an instance of the application
	a := app.New()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err = a.Server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
