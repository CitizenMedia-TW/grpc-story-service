package app

import (
	"context"
	"grpc-story-service/protobuffs/story-service"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (a *App) GetOneStory(ctx context.Context, in *story.GetOneStoryRequest) (*story.GetOneStoryResponse, error) {
	result, err := a.database.GetStoryById(ctx, in.StoryId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var comments = make([]*story.Comment, len(result.Comments))

	for i, c := range result.Comments {
		var subComments = make([]*story.SubComment, len(c.SubComments))
		for j, sc := range c.SubComments {
			subComments[j] = &story.SubComment{
				Id:               sc.Id,
				Content:          sc.Content,
				Commenter:        sc.CommenterName,
				CommenterId:      sc.CommenterId,
				Time:             timestamppb.New(sc.CreatedAt),
				RepliedCommentId: c.Id,
			}
		}
		comments[i] = &story.Comment{
			Id:          c.Id,
			Content:     c.Content,
			Commenter:   c.CommenterName,
			CommenterId: c.CommenterId,
			Time:        timestamppb.New(c.CreatedAt),
			SubComments: subComments,
		}
	}

	res := &story.GetOneStoryResponse{
		Author:    result.AuthorName,
		AuthorId:  result.AuthorId,
		Content:   result.Content,
		Comments:  comments,
		Title:     result.Title,
		SubTitle:  result.SubTitle,
		CreatedAt: timestamppb.New(result.CreatedAt),
		Tags:      result.Tags,
	}

	return res, nil
}
