package app

import (
	"google.golang.org/grpc"
	"grpc-story-service/proto"
	"grpc-story-service/internal/database"
)

type App struct {
	database *database.Database
	Server *grpc.Server
	story.UnimplementedStoryServiceServer
}

func New() *App {
	db := database.New()

	s := grpc.NewServer()
    story.RegisterStoryServiceServer(s, &App{database: db})

	return &App{
		database: db,
		Server: s,
	}
}
