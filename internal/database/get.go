package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"grpc-story-service/internal/models"
)

type StoryQuery struct {
	StoryEntity
	AuthorName string         `bson:"authorName"`
	Comments   []CommentQuery `bson:"comments"`
}

type CommentQuery struct {
	CommentEntity
	CommenterName string            `bson:"commenterName"`
	SubComments   []SubCommentQuery `bson:"subComments"`
}

type SubCommentQuery struct {
	SubCommentEntity
	ReplierName string `bson:"replierName"`
}

// GetStoryById todo: aggregate author name , commenter name
func (db *Database) GetStoryById(ctx context.Context, id string) (model.Story, error) {
	storyOid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Story{}, errors.Join(err, errors.New("invalid story id"+id))
	}

	pipeline := bson.A{
		bson.D{{"$match", bson.D{{"_id", storyOid}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Users"},
					{"localField", "authorId"},
					{"foreignField", "_id"},
					{"as", "authorName"},
				},
			},
		},
		bson.D{{"$set", bson.D{{"authorName", bson.D{{"$first", "$authorName.username"}}}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", CommentCollection},
					{"localField", "_id"},
					{"foreignField", "storyId"},
					{"as", "comments"},
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$comments"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Users"},
					{"localField", "comments.commenterId"},
					{"foreignField", "_id"},
					{"as", "comments.commenterName"},
				},
			},
		},
		bson.D{{"$set", bson.D{{"comments.commenterName", bson.D{{"$first", "$comments.commenterName.username"}}}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "StorySubComments"},
					{"localField", "comments._id"},
					{"foreignField", "parentId"},
					{"as", "comments.subComments"},
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$comments.subComments"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Users"},
					{"localField", "comments.subComments.commenterId"},
					{"foreignField", "_id"},
					{"as", "comments.subComments.commenterName"},
				},
			},
		},
		bson.D{{"$set", bson.D{{"comments.subComments.commenterName", bson.D{{"$first", "$comments.subComments.commenterName.username"}}}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "_id"},
					{"title", bson.D{{"$first", "$title"}}},
					{"authorName", bson.D{{"$first", "$authorName"}}},
					{"authorId", bson.D{{"$first", "$authorId"}}},
					{"comments", bson.D{{"$push", "$comments"}}},
				},
			},
		},
	}

	cursor, err := db.database.Collection(StoryCollection).Aggregate(ctx, pipeline)
	if err != nil {
		return model.Story{}, err
	}
	defer cursor.Close(ctx)
	cursor.Next(ctx)

	var storyQuery StoryQuery
	err = cursor.Decode(&storyQuery)
	if err != nil {
		return model.Story{}, err
	}

	return storyQuery.ToDomain(), nil
}

func (db *Database) GetStories(ctx context.Context, skip int32, count int32) ([]model.Story, error) {

	pipeline := bson.A{
		bson.D{{"$skip", skip}},
		bson.D{{"$limit", count}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Users"},
					{"localField", "authorId"},
					{"foreignField", "_id"},
					{"as", "authorName"},
				},
			},
		},
		bson.D{{"$set", bson.D{{"authorName", bson.D{{"$first", "$authorName.username"}}}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "StoryComments"},
					{"localField", "_id"},
					{"foreignField", "storyId"},
					{"as", "comments"},
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$comments"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Users"},
					{"localField", "comments.commenterId"},
					{"foreignField", "_id"},
					{"as", "comments.commenterName"},
				},
			},
		},
		bson.D{{"$set", bson.D{{"comments.commenterName", bson.D{{"$first", "$comments.commenterName.username"}}}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "StorySubComments"},
					{"localField", "comments._id"},
					{"foreignField", "parentId"},
					{"as", "comments.subComments"},
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$comments.subComments"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Users"},
					{"localField", "comments.subComments.commenterId"},
					{"foreignField", "_id"},
					{"as", "comments.subComments.commenterName"},
				},
			},
		},
		bson.D{{"$set", bson.D{{"comments.subComments.commenterName", bson.D{{"$first", "$comments.subComments.commenterName.username"}}}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$_id"},
					{"doc", bson.D{{"$first", "$$ROOT"}}},
				},
			},
		},
		bson.D{{"$replaceRoot", bson.D{{"newRoot", "$doc"}}}},
	}

	cursor, err := db.database.Collection(StoryCollection).Aggregate(ctx, pipeline)
	var stories []model.Story
	if err != nil {
		return stories, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var storyQuery StoryQuery
		err = cursor.Decode(&storyQuery)
		if err != nil {
			return stories, err
		}
		stories = append(stories, storyQuery.ToDomain())
	}

	return stories, nil
}

func (q StoryQuery) ToDomain() model.Story {
	comments := make([]model.Comment, len(q.Comments))
	for i, comment := range q.Comments {
		comments[i] = comment.ToDomain()
		comments[i].SubComments = make([]model.SubComment, len(comment.SubComments))
		for j, subComment := range comment.SubComments {
			comments[i].SubComments[j] = subComment.ToDomain()
		}
	}
	return model.Story{
		Id:        q.Id.Hex(),
		AuthorId:  q.AuthorId.Hex(),
		Content:   q.Content,
		Title:     q.Title,
		SubTitle:  q.SubTitle,
		CreatedAt: q.CreatedAt.Time(),
		Tags:      q.Tags,
		Comments:  comments,
	}
}
