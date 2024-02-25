package database

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete comment
func (db *Database) DeleteComment(s primitive.ObjectID, c primitive.ObjectID) error {
	return errors.ErrUnsupported
}

func (db *Database) DeleteStory(s primitive.ObjectID) error {
	return errors.ErrUnsupported
}

// DropCollection drops the collection (for testing purposes)
func (db *Database) DropCollection() error {
	err := db.database.Collection(StoryCollection).Drop(context.Background())
	if err != nil {
		log.Println("Error in DropCollection")
		return err
	}

	err = db.database.Collection(CommentCollection).Drop(context.Background())
	if err != nil {
		log.Println("Error in DropCollection")
		return err
	}

	err = db.database.Collection(SubCommentCollection).Drop(context.Background())
	if err != nil {
		log.Println("Error in DropCollection")
		return err
	}
	return nil
}
