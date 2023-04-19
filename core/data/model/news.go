package model

import (
	"time"
)

type News struct {
	Id            string
	Title         string
	Content       string
	ContentJp     string
	Description   string
	DescriptionJp string
	ImgUrl        string
	MetaKw        string
	MetaDesc      string
	Slug          string
	Category      int
	SubCategory   int
	CommentNum    int
	VoteNum       int
	ViewNum       int
	Status        int
	PublishBy     string
	Ranking       float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	PublishedAt   time.Time
}
