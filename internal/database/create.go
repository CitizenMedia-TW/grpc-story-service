package database

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewStory struct {
	AuthorId string
	Content  string
	Title    string
	SubTitle string
	Tags     []string
}

func (db *Database) NewStory(ctx context.Context, story NewStory) error {
	authorId, err := primitive.ObjectIDFromHex(story.AuthorId)
	if err != nil {
		return errors.Join(err, errors.New("invalid author id"+story.AuthorId))
	}

	//probably should check if author exist, but since it's nosql database, and it does not affect the query outcome, we'll skip it for now.
	storyEntity := StoryEntity{
		Id:        primitive.NewObjectID(),
		AuthorId:  authorId,
		Content:   story.Content,
		Title:     story.Title,
		SubTitle:  story.SubTitle,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		Tags:      story.Tags,
	}

	result, err := db.database.Collection(StoryCollection).InsertOne(ctx, storyEntity)

	if err != nil {
		log.Println("Error in push")
		return err
	}
	//todo: use a better logging/tracing system
	log.Print(result)
	return nil
}

func (db *Database) NewComment(ctx context.Context, commentedStoryId string, commenterId string, content string) error {
	storyOid, err := primitive.ObjectIDFromHex(commentedStoryId)
	if err != nil {
		return errors.Join(err, errors.New("invalid story id"+commentedStoryId))
	}
	commenterOid, err := primitive.ObjectIDFromHex(commenterId)
	if err != nil {
		return errors.Join(err, errors.New("invalid commenter id"+commenterId))
	}

	//probably should check if story and commenter exist, but since it's nosql database, and it does not affect the query outcome, we'll skip it for now.
	commentEntity := CommentEntity{
		Id:          primitive.NewObjectID(),
		StoryId:     storyOid,
		Content:     content,
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		CommenterId: commenterOid,
	}

	_, err = db.database.Collection(CommentCollection).InsertOne(ctx, commentEntity)

	if err != nil {
		return err
	}

	return nil
}

func (db *Database) NewSubComment(ctx context.Context, repliedCommentId string, replierId string, content string) error {
	repliedCommentOid, err := primitive.ObjectIDFromHex(repliedCommentId)
	if err != nil {
		return errors.Join(err, errors.New("invalid story id"+repliedCommentId))
	}
	replierOid, err := primitive.ObjectIDFromHex(replierId)
	if err != nil {
		return errors.Join(err, errors.New("invalid commenter id"+repliedCommentId))
	}

	reply := SubCommentEntity{
		Id:        primitive.NewObjectID(),
		ParentId:  repliedCommentOid,
		Content:   content,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		ReplierId: replierOid,
	}

	_, err = db.database.Collection(SubCommentCollection).InsertOne(ctx, reply)

	if err != nil {
		return err
	}

	return nil
}
