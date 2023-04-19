package model

type Category struct {
	Id int
	ParentId int
	Name string
	Description string
	Slug string
	NewsCount int
	Status int // 3: deleted
	Order int
}