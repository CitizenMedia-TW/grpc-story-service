package app

import (
	"context"
	"grpc-story-service/protobuffs/story-service"
	"log"
)

type GRPCServer interface {
	/* Test functions */
	SayHello(context.Context, *story.Empty) (*story.HelloReply, error)
	DropCollection(context.Context, *story.Empty) (*story.Empty, error)

	/* Create functions */
	CreateComment(context.Context, *story.CreateCommentRequest) (*story.CreateCommentResponse, error)
	CreateStory(context.Context, *story.CreateStoryRequest) (*story.CreateStoryResponse, error)

	/* Delete functions */
	DeleteComment(context.Context, *story.DeleteCommentRequest) (*story.DeleteCommentResponse, error)
	DeteteStory(context.Context, *story.DeleteStoryRequest) (*story.DeleteStoryResponse, error)

	/* Get functions */
	GetOneStory(context.Context, *story.GetOneStoryRequest) (*story.GetOneStoryResponse, error)
	GetRecommended(context.Context, *story.GetRecommendedRequest) (*story.GetRecommendedResponse, error)
}

func (a *App) SayHello(ctx context.Context, in *story.Empty) (*story.HelloReply, error) {
	return &story.HelloReply{Message: "greet !!!"}, nil
}

func (a *App) DropCollection(ctx context.Context, in *story.Empty) (*story.Empty, error) {
	err := a.database.DropCollection()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &story.Empty{}, nil
}

// Temperarily function
func (a *App) GetRecommended(ctx context.Context, in *story.GetRecommendedRequest) (*story.GetRecommendedResponse, error) {
	stories, err := a.database.GetStories(ctx, in.Skip, in.Count)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resList := make([]string, len(stories))
	for i, v := range stories {
		resList[i] = v.Id
	}

	return &story.GetRecommendedResponse{
		StoryIdList: resList,
	}, nil
}
