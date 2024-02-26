package app

import (
	"grpc-story-service/internal/database"
)

type App struct {
	database database.Database
}

func New() App {
	db := database.New()

	return App{
		database: db,
	}
}
