package outbound

import (
	"lykafe/news/core/data/model"
)

type CategoryRepo interface {
	GetAllCategories() ([]*model.Category, error)
	DeleteCategory(id int) (*model.Category, error) 
	UndeleteCategory(id int) (*model.Category, error)
	EditCategory(id int, name, slug string) (*model.Category, error)
	NewCategory(req *model.Category) (*model.Category, error)
	ReorderCategories(categories []*model.Category) ([]*model.Category, error)
}