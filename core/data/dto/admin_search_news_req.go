package dto

import "time"

type AdminSearchNewsReq struct {
	Title string
	Content string
	Category int
	SubCategory int
	Description string
	PublishBy string
	UpdateBy string
	Status int
	From time.Time
	To time.Time
	Offset int
	Limit int
}
