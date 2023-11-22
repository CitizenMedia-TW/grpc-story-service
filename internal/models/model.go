package model

import "go.mongodb.org/mongo-driver/bson/primitive"
import "time"

type Story struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Author    string             `bson:"author" json:"author"`
	AuthorId  primitive.ObjectID `bson:"authorId,omitempty" json:"authorId,omitempty"`
	Content   string             `bson:"content" json:"content"`
	Title     string             `bson:"title" json:"title"`
	SubTitle  string             `bson:"subTitle" json:"subTitle"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	Comments  []Comment          `bson:"comments,omitempty" json:"comments,omitempty"`
	Tags      []string           `bson:"tags" json:"tags"`
}

type Comment struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Comment     string             `bson:"comment" json:"comment"`
	Time        time.Time          `bson:"date" json:"date"`
	Commenter   string             `bson:"commenter" json:"commenter"`
	CommenterId primitive.ObjectID `bson:"commenterId" json:"commenterId"`
}
