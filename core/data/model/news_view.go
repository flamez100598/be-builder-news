package model

import (
	"time"
)

type NewsView struct {
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
	UpdateBy      string
	Ranking       float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	PublishedAt   time.Time
	// extra
	Order             int
	UsernamePublishBy string
	NamePublishBy     string
	AvatarPublishBy   string
	UsernameUpdateBy  string
	NameUpdateBy      string
	AvatarUpdateBy    string
}
