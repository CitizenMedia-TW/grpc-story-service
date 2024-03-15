package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"grpc-story-service/internal/models"
)

type StoryQuery struct {
	StoryEntity `bson:",inline"`
	AuthorName  string         `bson:"authorName"`
	Comments    []CommentQuery `bson:"comments"`
}

type CommentQuery struct {
	CommentEntity `bson:",inline"`
	CommenterName string            `bson:"commenterName"`
	SubComments   []SubCommentQuery `bson:"subComments"`
}

type SubCommentQuery struct {
	SubCommentEntity `bson:",inline"`
	ReplierName      string `bson:"replierName"`
}

// GetStoryById todo: aggregate author name , commenter name
func (db *Database) GetStoryById(ctx context.Context, id string) (model.Story, error) {
	storyId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Story{}, err
	}
	storyQuery, err := getStoryQuery(ctx, db.database, storyId)
	if err != nil {
		return model.Story{}, err
	}
	comments, err := getStoryComments(ctx, db.database, storyId)
	if err != nil {
		return model.Story{}, err
	}
	for i, comment := range comments {
		subComments, err := getSubComments(ctx, db.database, comment.Id)
		if err != nil {
			return model.Story{}, err
		}
		comments[i].SubComments = subComments
	}
	storyQuery.Comments = comments
	return storyQuery.ToDomain(), nil
}

func (db *Database) GetUserStoryId(ctx context.Context, id string) ([]string, error) {
	userId, err := primitive.ObjectIDFromHex(id)
	var storyIds []string
	if err != nil {
		return storyIds, err
	}
	cursor, err := db.database.Collection(StoryCollection).Find(ctx, bson.M{"authorId": userId})
	if err != nil {
		return storyIds, err
	}
	for cursor.Next(ctx) {
		doc := bson.M{}
		err = cursor.Decode(&doc)
		if err != nil {
			return storyIds, err
		}
		storyIds = append(storyIds, cursor.Current.Lookup("_id").ObjectID().String())
	}
	return storyIds, err
}

func (db *Database) GetStories(ctx context.Context, skip int64, count int64) ([]model.Story, error) {
	stories, err := getStories(ctx, db.database, skip, count)
	if err != nil {
		return nil, err
	}
	var results []model.Story
	for _, story := range stories {
		comments, err := getStoryComments(ctx, db.database, story.Id)
		if err != nil {
			return nil, err
		}
		for i, comment := range comments {
			subComments, err := getSubComments(ctx, db.database, comment.Id)
			if err != nil {
				return nil, err
			}
			comments[i].SubComments = subComments
		}
		story.Comments = comments
		results = append(results, story.ToDomain())
	}
	return results, nil
}

func getStories(ctx context.Context, database *mongo.Database, skip int64, count int64) ([]StoryQuery, error) {
	pipeline := bson.A{
		bson.D{{"$sort", bson.M{"_id": -1}}},
		bson.D{{"$skip", skip}},
		bson.D{{"$limit", count}},
		bson.D{{"$lookup", bson.D{{"from", "Users"}, {"localField", "authorId"}, {"foreignField", "_id"}, {"as", "authorName"}}}},
		bson.D{{"$set", bson.D{{"authorName", bson.D{{"$first", "$authorName.username"}}}}}},
	}
	cursor, err := database.Collection(StoryCollection).Aggregate(ctx, pipeline)
	var results []StoryQuery
	if err != nil {
		return results, err
	}
	for cursor.Next(ctx) {
		var entity StoryQuery
		err := cursor.Decode(&entity)
		if err != nil {
			return results, err
		}
		results = append(results, entity)
	}
	return results, nil
}

func getStoryQuery(ctx context.Context, database *mongo.Database, id primitive.ObjectID) (StoryQuery, error) {
	pipeline := bson.A{
		bson.D{{"$match", bson.D{{"_id", id}}}},
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
	}
	cursor, err := database.Collection(StoryCollection).Aggregate(ctx, pipeline)
	if err != nil {
		return StoryQuery{}, err
	}
	if !cursor.Next(ctx) {
		return StoryQuery{}, ErrNotFound
	}
	println(cursor.Current.String())
	entity := StoryQuery{}
	err = cursor.Decode(&entity)
	println(entity.AuthorName)
	return entity, err
}

func getStoryComments(ctx context.Context, database *mongo.Database, id primitive.ObjectID) ([]CommentQuery, error) {
	pipeline := bson.A{
		bson.D{{"$match", bson.D{{"storyId", id}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Users"},
					{"localField", "commenterId"},
					{"foreignField", "_id"},
					{"as", "commenterName"},
				},
			},
		},
		bson.D{{"$set", bson.D{{"commenterName", bson.D{{"$first", "$commenterName.username"}}}}}},
	}
	cursor, err := database.Collection(CommentCollection).Aggregate(ctx, pipeline)
	var results []CommentQuery
	if err != nil {
		return results, err
	}
	for cursor.Next(ctx) {
		var entity CommentQuery
		err := cursor.Decode(&entity)
		if err != nil {
			return results, err
		}
		results = append(results, entity)
	}
	return results, nil
}

func getSubComments(ctx context.Context, database *mongo.Database, id primitive.ObjectID) ([]SubCommentQuery, error) {
	pipeline := bson.A{
		bson.D{{"$match", bson.D{{"parentId", id}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Users"},
					{"localField", "replierId"},
					{"foreignField", "_id"},
					{"as", "replierName"},
				},
			},
		},
		bson.D{{"$set", bson.D{{"replierName", bson.D{{"$first", "$replierName.username"}}}}}},
	}
	cursor, err := database.Collection(SubCommentCollection).Aggregate(ctx, pipeline)
	var results []SubCommentQuery
	if err != nil {
		return results, err
	}
	for cursor.Next(ctx) {
		var entity SubCommentQuery
		err := cursor.Decode(&entity)
		if err != nil {
			return results, err
		}
		results = append(results, entity)
	}
	return results, nil
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
		Id:         q.Id.Hex(),
		AuthorId:   q.AuthorId.Hex(),
		AuthorName: q.AuthorName,
		Content:    q.Content,
		Title:      q.Title,
		SubTitle:   q.SubTitle,
		CreatedAt:  q.CreatedAt.Time(),
		Tags:       q.Tags,
		Comments:   comments,
	}
}
