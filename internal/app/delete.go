package app

import (
	"context"
	"grpc-story-service/protobuffs/story-service"
	"log"
)

func (a *App) DeleteComment(ctx context.Context, in *story.DeleteCommentRequest) (*story.DeleteCommentResponse, error) {

	err := a.database.DeleteComment(ctx, in.CommentId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &story.DeleteCommentResponse{Message: "success"}, nil
}

func (a *App) DeleteStory(ctx context.Context, in *story.DeleteStoryRequest) (*story.DeleteStoryResponse, error) {

	err := a.database.DeleteStory(ctx, in.StoryId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &story.DeleteStoryResponse{Message: "success"}, nil
}

func (a *App) DeleteSubComment(ctx context.Context, in *story.DeleteSubCommentRequest) (*story.DeleteSubCommentResponse, error) {

	err := a.database.DeleteSubComment(ctx, in.SubCommentId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &story.DeleteSubCommentResponse{Message: "success"}, nil
}
