package model

import (
	"time"
)

type Story struct {
	Id         string    `json:"id,omitempty"`
	AuthorName string    `json:"authorName"`
	AuthorId   string    `json:"authorId,omitempty"`
	Content    string    `json:"content"`
	Title      string    `json:"title"`
	SubTitle   string    `json:"subTitle"`
	CreatedAt  time.Time `json:"createdAt"`
	Tags       []string  `json:"tags"`
	Comments   []Comment `json:"comments"`
}

type Comment struct {
	Id            string       `json:"_id,omitempty"`
	Content       string       `json:"comment"`
	CreatedAt     time.Time    `json:"createdAt"`
	CommenterId   string       `json:"commenterId"`
	CommenterName string       `json:"commenterName"`
	SubComments   []SubComment `json:"replies"`
}

type SubComment struct {
	Id            string    `json:"_id,omitempty"`
	Content       string    `json:"comment"`
	CreatedAt     time.Time `json:"createdAt"`
	CommenterId   string    `json:"commenterId"`
	CommenterName string    `json:"commenterName"`
}
