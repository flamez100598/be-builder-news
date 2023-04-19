package usecase

import (
	"lykafe/news/core/data/model"
	"lykafe/news/core/data/dto"
)

type News interface {
	GetAllCategories() ([]*model.Category, error)
	DeleteCategory(id int) (*model.Category, error)
	UndeleteCategory(id int) (*model.Category, error)
	EditCategory(id int, name, slug string) (*model.Category, error)
	NewCategory(cat *model.Category) (*model.Category, error)
	ReorderCategories(categories []*model.Category) ([]*model.Category, error)
	PublishNews(req *dto.PublishNewsReq) (*model.News, error)
	EditNewsPost(req *dto.PublishNewsReq) (*model.News, error)
	DeleteNewsPost(id string) (*model.News, error)
	GetFeaturedNews() ([]*model.NewsView, error)
	GetNewsBySlug(slug string) (*model.NewsView, error)
	GetNewsById(id string) (*model.NewsView, error)
	GetNewsCategory(includeSubs bool, id, offset, limit int) ([]*model.NewsView, error)
	GetNewsSubCategory(id, offset, limit int) ([]*model.NewsView, error)
	AdminSearchNews(req *dto.AdminSearchNewsReq) ([]*model.NewsView, error)
	AdminSearchNewsCount(req *dto.AdminSearchNewsReq) (int, error)
	AdminSearchRelatedNews(req *dto.AdminSearchRelatedNewsReq) ([]*model.NewsView, error)
	AdminSearchRelatedNewsCount(req *dto.AdminSearchRelatedNewsReq) (int, error)
	SearchNews(req *dto.SearchNewsReq) ([]*model.NewsView, error)
	SearchNewsCount(req *dto.SearchNewsReq) (int, error)
	GetNewsByTag(tag string, offset, limit int, includeDeleted bool) ([]*model.NewsView, error)
	GetNewsByTagCount(tag string, includeDeleted bool) (int, error)
	GetRelatedNews(id string) ([]*model.NewsView, error)
	GetNewsTags(id string) ([]string, error) // get tags of a news
	GetTagsNews(tag string, offset, limit int) ([]*model.Tag, error) // get list tags
	GetTagsNewsCount(tag string) (int, error) // get list tags count
	ReorderFeatured(req []*dto.NewsOrder) ([]*dto.NewsOrder, error)
	GetAllNewsUrls() ([]*dto.NewsUrl, error)
}