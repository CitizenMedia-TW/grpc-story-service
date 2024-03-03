package app

import (
	"flag"
	"fmt"
	"grpc-story-service/internal/restapp"
	"grpc-story-service/protobuffs/auth-service"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func StartServer() {
	grpcClient, err := grpc.Dial("157.230.46.45:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	authClient := auth.NewAuthServiceClient(grpcClient)

	restServer := restapp.New(authClient)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = http.Serve(lis, restServer.Routes())
	if err != nil {
		return
	}
}
