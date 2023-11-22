package database

import (
	"context"
	"errors"
	"grpc-story-service/internal/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete comment
func (db *Database) DeleteComment(s primitive.ObjectID, c primitive.ObjectID) error {
	filter := bson.M{"_id": s}
	update := bson.M{
		"$pull": bson.M{
			"comments": bson.M{
				"_id": c,
			},
		},
	}

	result, err := db.database.UpdateOne(context.Background(), filter, update)
	if result.MatchedCount == 0 {
		log.Println("Commented story id not found")
		return errors.New("Commented story id not found")
	}
	if result.ModifiedCount == 0 {
		log.Println("Nothing modified, comment may be deleted or not exists")
		return errors.New("Nothing modified, comment may be deleted or not exists")
	}
	if err != nil {
		log.Println("Error in DeleteComment")
		return err
	}
	return nil
}

func (db *Database) DeleteStory(s primitive.ObjectID) error {
	var story *model.Story
	err := db.database.FindOneAndDelete(context.Background(), bson.D{primitive.E{Key: "_id", Value: s}}).Decode(&story)
	if err != nil {
		return err
	}
	return nil
}

// DropCollection drops the collection (for testing purposes)
func (db *Database) DropCollection() error {
	err := db.database.Drop(context.Background())
	if err != nil {
		log.Println("Error in DropCollection")
		return err
	}

	return nil
}
