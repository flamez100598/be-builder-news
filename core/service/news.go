package service

import (
	"lykafe/news/core/data/model"
	"lykafe/news/core/port/outbound"
	"lykafe/news/core/data/dto"
)

type NewsService struct {
	categoryRepo outbound.CategoryRepo
	newsRepo outbound.NewsRepo
}

func NewNewsService(categoryRepo outbound.CategoryRepo,
	newsRepo outbound.NewsRepo) *NewsService {
	return &NewsService{
		categoryRepo: categoryRepo,
		newsRepo: newsRepo,
	}
}

func (n *NewsService) GetAllCategories() ([]*model.Category, error) {
	return n.categoryRepo.GetAllCategories()
}

func (this *NewsService) DeleteCategory(id int) (*model.Category, error) {
	return this.categoryRepo.DeleteCategory(id)
}

func (this *NewsService) UndeleteCategory(id int) (*model.Category, error) {
	return this.categoryRepo.UndeleteCategory(id)
}

func (this *NewsService) EditCategory(id int, name, slug string) (*model.Category, error) {
	return this.categoryRepo.EditCategory(id, name, slug)
}

func (this *NewsService) NewCategory(cat *model.Category) (*model.Category, error) {
	return this.categoryRepo.NewCategory(cat)
}

func (this *NewsService) ReorderCategories(categories []*model.Category) ([]*model.Category, error) {
	return this.categoryRepo.ReorderCategories(categories)
}

func (n *NewsService) PublishNews(req *dto.PublishNewsReq) (*model.News, error) {
	return n.newsRepo.PublishNews(req)
}

func (n *NewsService) GetFeaturedNews() ([]*model.NewsView, error) {
	return n.newsRepo.GetFeaturedNews()
}

func (n *NewsService) GetNewsBySlug(slug string) (*model.NewsView, error) {
	return n.newsRepo.GetNewsBySlug(slug)
}

func (n *NewsService) GetNewsById(id string) (*model.NewsView, error) {
	return n.newsRepo.GetNewsById(id)
}

func (n *NewsService) GetNewsCategory(includeSubs bool, id, offset, limit int) ([]*model.NewsView, error) {
	return n.newsRepo.GetNewsCategory(includeSubs, id, offset, limit)
}

func (n *NewsService) GetNewsSubCategory(id, offset, limit int) ([]*model.NewsView, error) {
	return n.newsRepo.GetNewsSubCategory(id, offset, limit)
}

func (n *NewsService) AdminSearchNews(req *dto.AdminSearchNewsReq) ([]*model.NewsView, error) {
	return n.newsRepo.AdminSearchNews(req)
}

func (n *NewsService) AdminSearchNewsCount(req *dto.AdminSearchNewsReq) (int, error) {
	return n.newsRepo.AdminSearchNewsCount(req)
}

func (this *NewsService) AdminSearchRelatedNews(req *dto.AdminSearchRelatedNewsReq) ([]*model.NewsView, error) {
	return this.newsRepo.AdminSearchRelatedNews(req)
}

func (this *NewsService) AdminSearchRelatedNewsCount(req *dto.AdminSearchRelatedNewsReq) (int, error) {
	return this.newsRepo.AdminSearchRelatedNewsCount(req)
}

func (n *NewsService) SearchNews(req *dto.SearchNewsReq) ([]*model.NewsView, error) {
	return n.newsRepo.SearchNews(req)
}

func (n *NewsService) SearchNewsCount(req *dto.SearchNewsReq) (int, error) {
	return n.newsRepo.SearchNewsCount(req)
}

func (n *NewsService) DeleteNewsPost(id string) (*model.News, error) {
	return n.newsRepo.DeleteNewsPost(id)
}

func (n *NewsService) EditNewsPost(req *dto.PublishNewsReq) (*model.News, error) {
	return n.newsRepo.EditNewsPost(req)
}

func (n *NewsService) GetNewsByTag(tag string, offset, limit int, includeDeleted bool) ([]*model.NewsView, error) {
	return n.newsRepo.GetNewsByTag(tag, offset, limit, includeDeleted)
}

func (n *NewsService) GetNewsByTagCount(tag string, includeDeleted bool) (int, error) {
	return n.newsRepo.GetNewsByTagCount(tag, includeDeleted)
}

func (n *NewsService) GetRelatedNews(id string) ([]*model.NewsView, error) {
	return n.newsRepo.GetRelatedNews(id)
}

func (n *NewsService) GetNewsTags(id string) ([]string, error) {
	return n.newsRepo.GetNewsTags(id)
}

// get list tags
func (n *NewsService) GetTagsNews(tag string, offset, limit int) ([]*model.Tag, error) {
	return n.newsRepo.GetTagsNews(tag, offset, limit)
}

// get list tags count
func (n *NewsService) GetTagsNewsCount(tag string) (int, error) {
	return n.newsRepo.GetTagsNewsCount(tag)
}

func (n *NewsService) ReorderFeatured(req []*dto.NewsOrder) ([]*dto.NewsOrder, error) {
	return n.newsRepo.ReorderFeatured(req)
}

func (n *NewsService) GetAllNewsUrls() ([]*dto.NewsUrl, error) {
	return n.newsRepo.GetAllNewsUrls()
}