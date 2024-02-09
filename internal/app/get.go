package app

import (
	"context"
	"grpc-story-service/protobuffs/story-service"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (a *App) GetOneStory(ctx context.Context, in *story.GetOneStoryRequest) (*story.GetOneStoryResponse, error) {
	storyId, err := primitive.ObjectIDFromHex(in.StoryId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := a.database.GetStoryById(storyId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	comments := make([]*story.Comment, len(result.Comments))
	for i, comment := range result.Comments {
		comments[i] = &story.Comment{
			Id:          comment.Id.Hex(),
			Comment:     comment.Comment,
			Time:        timestamppb.New(comment.Time),
			Commenter:   comment.Commenter,
			CommenterId: comment.CommenterId.Hex(),
		}
	}

	res := &story.GetOneStoryResponse{
		Author:    result.Author,
		AuthorId:  result.AuthorId.Hex(),
		Comments:  comments,
		Title:     result.Title,
		SubTitle:  result.SubTitle,
		CreatedAt: timestamppb.New(result.CreatedAt),
		Tags:      result.Tags,
	}

	return res, nil
}
