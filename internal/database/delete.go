package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func (db *Database) DeleteComment(ctx context.Context, storyId string) error {
	result := db.database.Collection(CommentCollection).FindOneAndDelete(ctx, bson.M{"_id": storyId})
	if result.Err() != nil {
		log.Println("Error in DeleteComment")
		return result.Err()
	}
	comment := CommentEntity{}
	err := result.Decode(&comment)
	if err != nil {
		log.Println("Error in DeleteComment")
		return err
	}
	_, err = db.database.Collection(SubCommentCollection).DeleteMany(ctx, bson.M{"parentId": comment.Id})
	if err != nil {
		log.Println("Error in DeleteComment")
		return err
	}
	return nil
}

func (db *Database) DeleteSubComment(ctx context.Context, subCommentId string) error {
	_, err := db.database.Collection(SubCommentCollection).DeleteOne(ctx, bson.M{"_id": subCommentId})
	if err != nil {
		log.Println("Error in DeleteSubComment")
		return err
	}
	return nil
}

func (db *Database) DeleteStory(ctx context.Context, storyId string) error {
	_, err := db.database.Collection(StoryCollection).DeleteOne(ctx, bson.M{"_id": storyId})
	if err != nil {
		log.Println("Error in DeleteStory")
		return err
	}
	cursor, err := db.database.Collection(CommentCollection).Find(ctx, bson.M{"storyId": storyId}, options.Find())
	if err != nil {
		log.Println("Error in DeleteStory")
		return err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var commentEntity CommentEntity
		err = cursor.Decode(&commentEntity)
		if err != nil {
			log.Println("Error in DeleteStory")
			return err
		}
		_, err = db.database.Collection(SubCommentCollection).DeleteMany(ctx, bson.M{"parentId": commentEntity.Id})
		if err != nil {
			log.Println("Error in DeleteStory")
			return err
		}
	}
	_, err = db.database.Collection(CommentCollection).DeleteMany(ctx, bson.M{"storyId": storyId})
	if err != nil {
		log.Println("Error in DeleteStory")
		return err
	}
	return nil
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
