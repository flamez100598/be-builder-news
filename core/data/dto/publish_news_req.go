package dto

import "time"

type PublishNewsReq struct {
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
	PublishBy     string
	PublishAt     time.Time
	Category      int
	SubCategory   int
	IsFeatured    bool
	IsHot         bool
	IsTrending    bool
	Related       []string
	Tags          []string
	AuthorId      string
}
