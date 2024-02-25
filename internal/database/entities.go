package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	model "grpc-story-service/internal/models"
)

var StoryCollection = "Stories"

type StoryEntity struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId  primitive.ObjectID `bson:"authorId,omitempty"`
	Content   string             `bson:"content"`
	Title     string             `bson:"title"`
	SubTitle  string             `bson:"subTitle"`
	CreatedAt primitive.DateTime `bson:"createdAt"`
	Tags      []string           `bson:"tags"`
}

var CommentCollection = "StoryComments"

type CommentEntity struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	StoryId     primitive.ObjectID `bson:"storyId"`
	Content     string             `bson:"content"`
	CreatedAt   primitive.DateTime `bson:"createdAt"`
	CommenterId primitive.ObjectID `bson:"commenterId"`
}

var SubCommentCollection = "StorySubComments"

type SubCommentEntity struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	ParentId  primitive.ObjectID `bson:"parentId"`
	Content   string             `bson:"content"`
	CreatedAt primitive.DateTime `bson:"createdAt"`
	ReplierId primitive.ObjectID `bson:"replierId"`
}

func (e CommentEntity) ToDomain() model.Comment {
	return model.Comment{
		Id:          e.Id.Hex(),
		Content:     e.Content,
		CreatedAt:   e.CreatedAt.Time(),
		CommenterId: e.CommenterId.Hex(),
	}
}

func (e SubCommentEntity) ToDomain() model.SubComment {
	return model.SubComment{
		Id:          e.Id.Hex(),
		Content:     e.Content,
		CreatedAt:   e.CreatedAt.Time(),
		CommenterId: e.ReplierId.Hex(),
	}
}

func (e StoryEntity) ToDomain() model.Story {
	return model.Story{
		Id:        e.Id.Hex(),
		AuthorId:  e.AuthorId.Hex(),
		Content:   e.Content,
		Title:     e.Title,
		SubTitle:  e.SubTitle,
		CreatedAt: e.CreatedAt.Time(),
		Tags:      e.Tags,
	}
}
