package app

import (
	"context"
	"grpc-story-service/internal/database"
	"grpc-story-service/protobuffs/story-service"
)

func (a *App) CreateComment(ctx context.Context, in *story.CreateCommentRequest) (*story.CreateCommentResponse, error) {
	err := a.database.NewComment(ctx, in.CommentedStoryId, in.CommenterId, in.Comment)
	if err != nil {
		return nil, err
	}
	return &story.CreateCommentResponse{Message: "Success"}, nil
}

func (a *App) CreateStory(ctx context.Context, in *story.CreateStoryRequest) (*story.CreateStoryResponse, error) {
	err := a.database.NewStory(ctx, database.NewStory{
		AuthorId: in.AuthorId,
		Content:  in.Content,
		Title:    in.Title,
		SubTitle: in.SubTitle,
		Tags:     in.Tags,
	})

	if err != nil {
		return nil, err
	}

	return &story.CreateStoryResponse{Message: "Success"}, nil
}

func (a *App) CreateCommentReply(ctx context.Context, in *story.CreateSubCommentRequest) (*story.CreateSubCommentResponse, error) {
	err := a.database.NewSubComment(ctx, in.RepliedCommentId, in.CommenterId, in.Content)

	if err != nil {
		return nil, err
	}
	return &story.CreateSubCommentResponse{Message: "Success"}, nil
}
