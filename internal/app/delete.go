package app

import (
	"context"
	"grpc-story-service/protobuffs/story-service"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *App) DeleteComment(ctx context.Context, in *story.DeleteCommentRequest) (*story.DeleteCommentResponse, error) {
	deleteFrom, err := primitive.ObjectIDFromHex(in.StoryId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	deleteComment, err := primitive.ObjectIDFromHex(in.CommentId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = a.database.DeleteComment(deleteFrom, deleteComment)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &story.DeleteCommentResponse{Message: "success"}, nil
}

func (a *App) DeleteStory(ctx context.Context, in *story.DeleteStoryRequest) (*story.DeleteStoryResponse, error) {
	deleteStory, err := primitive.ObjectIDFromHex(in.StoryId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = a.database.DeleteStory(deleteStory)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &story.DeleteStoryResponse{Message: "success"}, nil
}
