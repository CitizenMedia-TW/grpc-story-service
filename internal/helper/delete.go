package helper

import (
	"context"
	"grpc-story-service/protobuffs/story-service"
	"log"
)

func (h *Helper) DeleteComment(ctx context.Context, in *story.DeleteCommentRequest) (*story.DeleteCommentResponse, error) {

	err := h.database.DeleteComment(ctx, in.CommentId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &story.DeleteCommentResponse{Message: "success"}, nil
}

func (h *Helper) DeleteStory(ctx context.Context, in *story.DeleteStoryRequest) (*story.DeleteStoryResponse, error) {

	err := h.database.DeleteStory(ctx, in.StoryId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &story.DeleteStoryResponse{Message: "success"}, nil
}

func (h *Helper) DeleteSubComment(ctx context.Context, in *story.DeleteSubCommentRequest) (*story.DeleteSubCommentResponse, error) {

	err := h.database.DeleteSubComment(ctx, in.SubCommentId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &story.DeleteSubCommentResponse{Message: "success"}, nil
}
