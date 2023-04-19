package dto

import "time"

type PublishNewsReq struct {
  Id string
	Title string
  Content string
  Description string
  ImgUrl string
  MetaKw string
  MetaDesc string
  Slug string
  PublishBy string
  PublishAt time.Time
  Category int
  SubCategory int
  IsFeatured bool
  IsHot bool
  IsTrending bool
  Related []string
  Tags []string
  AuthorId string
}