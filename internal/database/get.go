package database

import (
	"context"
	"grpc-story-service/internal/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *Database) GetStoryById(id primitive.ObjectID) (*model.Story, error) {
	var story model.Story
	err := db.database.FindOne(context.Background(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&story)
	if err != nil {
		log.Println("Error in GetStoryById")
		return nil, err
	}

	return &story, nil
}

func (db *Database) GetTen() ([]*model.Story, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	opts := options.Find().SetProjection(bson.D{{Key: "id", Value: 1}})

	cursor, err := db.database.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("Error in GetTen")
		return nil, err
	}

	defer cursor.Close(ctx)

	var stories []*model.Story
	for cursor.Next(ctx) {
		var item model.Story

		err := cursor.Decode(&item)
		if err != nil {
			log.Print("Error decoding log into slice:", err)
			return nil, err
		} else {
			stories = append(stories, &item)
		}

	}

	return stories, nil
}
