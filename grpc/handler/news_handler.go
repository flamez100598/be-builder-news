package handler

import (
	"context"
	pb "lykafe/news/grpc/news/pb"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"lykafe/news/core/port/usecase"
	"lykafe/news/grpc/mapper"
	"lykafe/news/core/data/codes"
	"lykafe/news/core/data/dto"
	"lykafe/news/core/data/model"
)

type NewsHandler struct {
	pb.UnimplementedNewsServiceServer
	newsUsecase usecase.News
}

func NewNewsHandler(newsUsecase usecase.News) *NewsHandler{
	return &NewsHandler{
		newsUsecase: newsUsecase,
	}
}

func (h *NewsHandler) GetAllCategories(ctx context.Context, in *pb.NewsEmpty) (*pb.CategoriesResp, error) {
	categories, err := h.newsUsecase.GetAllCategories()
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbCategories := []*pb.Category{}
	for _, cat := range categories {
		pbCategories = append(pbCategories, mapper.Model2PbCategory(cat))
	}
	return  &pb.CategoriesResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Categories: pbCategories,
	}, nil
}

func (h *NewsHandler) DeleteCategory(ctx context.Context, in *pb.NewsIdInt) (*pb.CategoryResp, error) {
	category, err := h.newsUsecase.DeleteCategory(int(in.Id))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.CategoryResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Category: &pb.Category{
			Id: int32(category.Id),
		},
	}, nil
}

func (h *NewsHandler) UndeleteCategory(ctx context.Context, in *pb.NewsIdInt) (*pb.CategoryResp, error) {
	category, err := h.newsUsecase.UndeleteCategory(int(in.Id))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.CategoryResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Category: &pb.Category{
			Id: int32(category.Id),
		},
	}, nil
}

func (h *NewsHandler) EditCategory(ctx context.Context, in *pb.Category) (*pb.CategoryResp, error) {
	category, err := h.newsUsecase.EditCategory(int(in.Id), in.Name, in.Slug)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.CategoryResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Category: &pb.Category{
			Id: int32(category.Id),
			Name: category.Name,
			Slug: category.Slug,
		},
	}, nil
}

func (h *NewsHandler) NewCategory(ctx context.Context, in *pb.Category) (*pb.CategoryResp, error) {
	category, err := h.newsUsecase.NewCategory(mapper.Pb2ModelCategory(in))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.CategoryResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Category: mapper.Model2PbCategory(category),
	}, nil
}

func (h *NewsHandler) ReorderCategories (ctx context.Context, in *pb.CategoriesReq) (*pb.CategoriesResp, error) {
	categories := []*model.Category{}
	for _, cate := range in.Categories {
		category := mapper.Pb2ModelCategory(cate)
		categories = append(categories, category)
	}
	cates, err := h.newsUsecase.ReorderCategories(categories)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbCategories := []*pb.Category{}
	for _, cate := range cates {
		pbCate := mapper.Model2PbCategory(cate)
		pbCategories = append(pbCategories, pbCate)
	}
	return  &pb.CategoriesResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Categories: pbCategories,
	}, nil
}

func (h *NewsHandler) PublishNews(ctx context.Context, in *pb.PublishNewsReq) (*pb.NewsResp, error) {
	req := mapper.Pb2DtoPublishNewsReq(in)
	news, err := h.newsUsecase.PublishNews(req)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.NewsResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		News: mapper.Model2PbNews(news),
	}, nil
}

func (h *NewsHandler) GetFeaturedNews(ctx context.Context, in *pb.NewsEmpty) (*pb.ListNewsViewResp, error) {
	listNews, err := h.newsUsecase.GetFeaturedNews()
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbNewsView := []*pb.NewsView{}
	for _, n := range listNews {
		pbNewsView = append(pbNewsView, mapper.Model2PbNewsView(n))
	}
	return  &pb.ListNewsViewResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		News: pbNewsView,
	}, nil
}

func (h *NewsHandler) GetNewsBySlug(ctx context.Context, in *pb.Slug) (*pb.NewsViewResp, error) {
	news, err := h.newsUsecase.GetNewsBySlug(in.Slug)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.NewsViewResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		News: mapper.Model2PbNewsView(news),
	}, nil
}

func (h *NewsHandler) GetNewsById(ctx context.Context, in *pb.NewsId) (*pb.NewsViewResp, error) {
	news, err := h.newsUsecase.GetNewsById(in.Id)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.NewsViewResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		News: mapper.Model2PbNewsView(news),
	}, nil
}

func (h *NewsHandler) GetNewsCategory(ctx context.Context, in *pb.NewsCategoryReq) (*pb.ListNewsViewResp, error) {
	listNews, err := h.newsUsecase.GetNewsCategory(in.IncludeSubs, int(in.CategoryId), int(in.Offset), int(in.Limit))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbNewsView := []*pb.NewsView{}
	for _, n := range listNews {
		pbNewsView = append(pbNewsView, mapper.Model2PbNewsView(n))
	}
	return  &pb.ListNewsViewResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		News: pbNewsView,
	}, nil
}

func (h *NewsHandler) GetNewsSubCategory(ctx context.Context, in *pb.NewsCategoryReq) (*pb.ListNewsViewResp, error) {
	listNews, err := h.newsUsecase.GetNewsSubCategory(int(in.CategoryId), int(in.Offset), int(in.Limit))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbNewsView := []*pb.NewsView{}
	for _, n := range listNews {
		pbNewsView = append(pbNewsView, mapper.Model2PbNewsView(n))
	}
	return  &pb.ListNewsViewResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		News: pbNewsView,
	}, nil
}

func (h *NewsHandler) AdminSearchNews(ctx context.Context, in *pb.AdminSearchNewsReq) (*pb.ListNewsViewResp, error) {
	listNews, err := h.newsUsecase.AdminSearchNews(mapper.Pb2DtoAdminSearchNewsReq(in))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbNewsView := []*pb.NewsView{}
	for _, n := range listNews {
		pbNewsView = append(pbNewsView, mapper.Model2PbNewsView(n))
	}
	return  &pb.ListNewsViewResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		News: pbNewsView,
	}, nil
}

func (h *NewsHandler) AdminSearchNewsCount(ctx context.Context, in *pb.AdminSearchNewsReq) (*pb.NewsCountResp, error) {
	c, err := h.newsUsecase.AdminSearchNewsCount(mapper.Pb2DtoAdminSearchNewsReq(in))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.NewsCountResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Count: int32(c),
	}, nil
}

func (h *NewsHandler) AdminSearchRelatedNews(ctx context.Context, in *pb.AdminSearchRelatedNewsReq) (*pb.ListNewsViewResp, error) {
	listNews, err := h.newsUsecase.AdminSearchRelatedNews(mapper.Pb2DtoAdminSearchRelatedNewsReq(in))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbNewsView := []*pb.NewsView{}
	for _, n := range listNews {
		pbNewsView = append(pbNewsView, mapper.Model2PbNewsView(n))
	}
	return  &pb.ListNewsViewResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		News: pbNewsView,
	}, nil
}

func (h *NewsHandler) AdminSearchRelatedNewsCount(ctx context.Context, in *pb.AdminSearchRelatedNewsReq) (*pb.NewsCountResp, error) {
	c, err := h.newsUsecase.AdminSearchRelatedNewsCount(mapper.Pb2DtoAdminSearchRelatedNewsReq(in))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.NewsCountResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Count: int32(c),
	}, nil
}

func (h *NewsHandler) SearchNews(ctx context.Context, in *pb.SearchNewsReq) (*pb.ListNewsViewResp, error) {
	listNews, err := h.newsUsecase.SearchNews(mapper.Pb2DtoSearchNewsReq(in))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbNewsView := []*pb.NewsView{}
	for _, n := range listNews {
		pbNewsView = append(pbNewsView, mapper.Model2PbNewsView(n))
	}
	return  &pb.ListNewsViewResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		News: pbNewsView,
	}, nil
}

func (h *NewsHandler) SearchNewsCount(ctx context.Context, in *pb.SearchNewsReq) (*pb.NewsCountResp, error) {
	c, err := h.newsUsecase.SearchNewsCount(mapper.Pb2DtoSearchNewsReq(in))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.NewsCountResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Count: int32(c),
	}, nil
}

func (h *NewsHandler) DeleteNewsPost (ctx context.Context, in *pb.NewsId) (*pb.NewsResp, error) {
	news, err := h.newsUsecase.DeleteNewsPost(in.Id)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.NewsResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		News: &pb.News{
			Id: news.Id,
		},
	}, nil
}

func (h *NewsHandler) EditNews(ctx context.Context, in *pb.PublishNewsReq) (*pb.NewsResp, error) {
	req := mapper.Pb2DtoPublishNewsReq(in)
	news, err := h.newsUsecase.EditNewsPost(req)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.NewsResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		News: mapper.Model2PbNews(news),
	}, nil
}

func (h *NewsHandler) GetNewsByTag(ctx context.Context, in *pb.NewsTagReq) (*pb.ListNewsViewResp, error) {
	listNews, err := h.newsUsecase.GetNewsByTag(in.Tag, int(in.Offset), int(in.Limit), in.IncludeDeleted)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbNewsView := []*pb.NewsView{}
	for _, n := range listNews {
		pbNewsView = append(pbNewsView, mapper.Model2PbNewsView(n))
	}
	return  &pb.ListNewsViewResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		News: pbNewsView,
	}, nil
}

func (h *NewsHandler) GetNewsByTagCount(ctx context.Context, in *pb.NewsTagReq) (*pb.NewsCountResp, error) {
	c, err := h.newsUsecase.GetNewsByTagCount(in.Tag, in.IncludeDeleted)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.NewsCountResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Count: int32(c),
	}, nil
}

func (h *NewsHandler) GetRelatedNews(ctx context.Context, in *pb.NewsId) (*pb.ListNewsViewResp, error) {
	listNews, err := h.newsUsecase.GetRelatedNews(in.Id)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbNewsView := []*pb.NewsView{}
	for _, n := range listNews {
		pbNewsView = append(pbNewsView, mapper.Model2PbNewsView(n))
	}
	return  &pb.ListNewsViewResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		News: pbNewsView,
	}, nil
}

func (h *NewsHandler) GetNewsTags (ctx context.Context, in *pb.NewsId) (*pb.NewsTagsResp, error) {
	tags, err := h.newsUsecase.GetNewsTags(in.Id)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.NewsTagsResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Tags: tags,
	}, nil
}

// get list tags
func (h *NewsHandler) GetTagsNews (ctx context.Context, in *pb.NewsTagReq) (*pb.TagNewsResp, error) {
	tags, err := h.newsUsecase.GetTagsNews(in.Tag, int(in.Offset), int(in.Limit))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbTags := []*pb.TagNews{}
	for _, t := range tags {
		pbTag := pb.TagNews{
			Id: t.Id,
			Name: t.Name,
			Count: int32(t.Count),
		}
		pbTags = append(pbTags, &pbTag)
	}
	return  &pb.TagNewsResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Tags: pbTags,
	}, nil
} 

// get list tags count
func (h *NewsHandler) GetTagsNewsCount (ctx context.Context, in *pb.NewsTagReq) (*pb.NewsCountResp, error) {
	c, err := h.newsUsecase.GetTagsNewsCount(in.Tag)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.NewsCountResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Count: int32(c),
	}, nil
}

func (h *NewsHandler) ReorderFeatured (ctx context.Context, in *pb.NewsOrders) (*pb.NewsOrdersResp, error) {
	req := []*dto.NewsOrder{}
	for _, o := range in.Orders {
		nO := dto.NewsOrder {
			Id: o.Id,
			Order: int(o.Order),
		}
		req = append(req, &nO)
	}
	orders, err := h.newsUsecase.ReorderFeatured(req)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbOrders := []*pb.NewsOrder{}
	for _, o := range orders {
		pbO := pb.NewsOrder{
			Id: o.Id,
			Order: int32(o.Order),
		}
		pbOrders = append(pbOrders, &pbO)
	}
	return  &pb.NewsOrdersResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Orders: pbOrders,
	}, nil
}

func (h *NewsHandler) GetAllNewsUrls (ctx context.Context, in *pb.NewsEmpty) (*pb.NewsUrlsResp, error) {
	urls, err := h.newsUsecase.GetAllNewsUrls()
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbNewsUrls := []*pb.NewsUrl{}
	for _, u := range urls {
		pbU := pb.NewsUrl{
			Id: u.Id,
			NewsSlug: u.NewsSlug,
			CategorySlug: u.CategorySlug,
		}
		pbNewsUrls = append(pbNewsUrls, &pbU)
	}
	return  &pb.NewsUrlsResp{
		Code: codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Urls: pbNewsUrls,
	}, nil
}