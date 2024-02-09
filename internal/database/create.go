package database

import (
	"context"
	"errors"
	"grpc-story-service/internal/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// New story
func (db *Database) NewStory(m model.Story) (*mongo.InsertOneResult, error) {
	result, err := db.database.InsertOne(context.Background(), m)
	if err != nil {
		log.Println("Error in push")
		return result, err
	}
	log.Print(result.InsertedID)
	return result, nil
}

// New Comment
func (db *Database) NewComment(m model.Comment, commentedStoryId primitive.ObjectID) error {
	filter := bson.M{"_id": commentedStoryId}
	update := bson.D{
		primitive.E{
			Key: "$push", Value: bson.D{primitive.E{Key: "comments", Value: m}},
		},
	}

	result, err := db.database.UpdateOne(context.Background(), filter, update)
	log.Println(result)
	if result.MatchedCount == 0 {
		log.Println("Commented story id not found")
		return errors.New("Commented story id not found")
	}
	if err != nil {
		log.Println("Error in CreateComment")
		return err
	}
	return nil
}
