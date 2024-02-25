package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	database *mongo.Database
}

func New() *Database {
	db := connect()
	return &Database{
		database: db,
	}
}

func connect() *mongo.Database {
	log.Println("Connecting to MongoDB...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use this for when running in Docker
	// clientOptions := options.Client().ApplyURI("mongodb://root:rootpassword@mongo:27017/")
	// client, err := mongo.Connect(ctx, clientOptions)

	// Use this for when running locally
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	return client.Database("GolangStoryTest")
}
