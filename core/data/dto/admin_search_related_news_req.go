package dto

import "time"

type AdminSearchRelatedNewsReq struct {
	Tags []string
	Title string
	Content string
	Category int
	SubCategory int
	PublishBy string
	From time.Time
	To time.Time
	Offset int
	Limit int
}
