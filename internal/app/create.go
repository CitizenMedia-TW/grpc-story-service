package app

import (
	"context"
	"grpc-story-service/internal/models"
	"grpc-story-service/proto"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *App) CreateComment(ctx context.Context, in *story.CreateCommentRequest) (*story.CreateCommentResponse, error) {
	// Convert author id
	commenterId, err := primitive.ObjectIDFromHex(in.CommenterId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	commentedStoryId, err := primitive.ObjectIDFromHex(in.CommentedStoryId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	newComment := model.Comment{
		Id:          primitive.NewObjectID(),
		Comment:     in.Comment,
		Time:        time.Now(),
		Commenter:   in.Commenter,
		CommenterId: commenterId,
	}

	err = a.database.NewComment(newComment, commentedStoryId)
	if err != nil {
		return nil, err
	}
	return &story.CreateCommentResponse{Message: "Success"}, nil
}

func (a *App) CreateStory(ctx context.Context, in *story.CreateStoryRequest) (*story.CreateStoryResponse, error) {
	// Convert author id
	authorId, err := primitive.ObjectIDFromHex(in.AuthorId)
	if err != nil {
		log.Println(err)
	}

	newStory := model.Story{
		Author:    in.Author,
		AuthorId:  authorId,
		Content:   in.Content,
		Title:     in.Title,
		SubTitle:  in.SubTitle,
		CreatedAt: time.Now(),
		Tags:      in.Tags,
		Comments:  []model.Comment{}, // Initialize to empty comment list
	}

	err = a.database.NewStory(newStory)
	if err != nil {
		return nil, err
	}
	return &story.CreateStoryResponse{Message: "Success"}, nil
}
