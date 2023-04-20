package mapper

import (
	"lykafe/news/core/data/dto"
	"lykafe/news/core/data/model"
	"lykafe/news/grpc/news/pb"
	"time"
	// "fmt"
)

func Pb2DtoPublishNewsReq(req *pb.PublishNewsReq) *dto.PublishNewsReq {
	if req == nil {
		return nil
	}
	publishAt, _ := time.Parse("2006-01-02 15:04:05", req.PublishAt)
	return &dto.PublishNewsReq{
		Id:            req.Id,
		Title:         req.Title,
		Content:       req.Content,
		ContentJp:     req.ContentJp,
		Description:   req.Description,
		DescriptionJp: req.DescriptionJp,
		ImgUrl:        req.ImgUrl,
		Slug:          req.Slug,
		MetaKw:        req.MetaKw,
		MetaDesc:      req.MetaDesc,
		IsFeatured:    req.IsFeatured,
		IsHot:         req.IsHot,
		IsTrending:    req.IsTrending,
		Tags:          req.Tags,
		Related:       req.Related,
		Category:      int(req.Category),
		SubCategory:   int(req.SubCategory),
		PublishBy:     req.PublishBy,
		PublishAt:     publishAt,
		AuthorId:      req.AuthorId,
	}
}

func Model2PbNews(news *model.News) *pb.News {
	if news == nil {
		return nil
	}
	return &pb.News{
		Title:         news.Title,
		Content:       news.Content,
		ContentJp:     news.ContentJp,
		Description:   news.Description,
		DescriptionJp: news.DescriptionJp,
		ImgUrl:        news.ImgUrl,
		MetaKw:        news.MetaKw,
		MetaDesc:      news.MetaDesc,
		Slug:          "",
		// IsFeatured: news.IsFeatured,
		// IsHot: news.IsHot,
		// IsTrending: news.IsTrending,
		// Tags: news.Tags,
		// Related: news.Related,
		Category:    int32(news.Category),
		SubCategory: int32(news.SubCategory),
		PublishBy:   news.PublishBy,
	}
}

func Model2PbNewsView(news *model.NewsView) *pb.NewsView {
	if news == nil {
		return nil
	}
	return &pb.NewsView{
		Id:                news.Id,
		Title:             news.Title,
		Content:           news.Content,
		ContentJp:         news.ContentJp,
		Description:       news.Description,
		DescriptionJp:     news.DescriptionJp,
		ImgUrl:            news.ImgUrl,
		MetaKw:            news.MetaKw,
		MetaDesc:          news.MetaDesc,
		Slug:              news.Slug,
		Category:          int32(news.Category),
		SubCategory:       int32(news.SubCategory),
		CommentNum:        int32(news.CommentNum),
		VoteNum:           int32(news.VoteNum),
		ViewNum:           int32(news.ViewNum),
		Status:            int32(news.Status),
		PublishBy:         news.PublishBy,
		Ranking:           news.Ranking,
		CreatedAt:         news.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:         news.UpdatedAt.Format("2006-01-02 15:04:05"),
		PublishedAt:       news.PublishedAt.Format("2006-01-02 15:04:05"),
		Order:             int32(news.Order),
		UsernamePublishBy: news.UsernamePublishBy,
		NamePublishBy:     news.NamePublishBy,
		AvatarPublishBy:   news.AvatarPublishBy,
		UsernameUpdateBy:  news.UsernameUpdateBy,
		NameUpdateBy:      news.NameUpdateBy,
		AvatarUpdateBy:    news.AvatarUpdateBy,
	}
}

func Pb2DtoAdminSearchNewsReq(req *pb.AdminSearchNewsReq) *dto.AdminSearchNewsReq {
	if req == nil {
		return nil
	}
	from, _ := time.Parse("2006-01-02 15:04:05", req.From)
	to, _ := time.Parse("2006-01-02 15:04:05", req.To)
	return &dto.AdminSearchNewsReq{
		Title:       req.Title,
		Content:     req.Content,
		Category:    int(req.Category),
		SubCategory: int(req.SubCategory),
		Description: req.Description,
		PublishBy:   req.PublishBy,
		UpdateBy:    req.UpdateBy,
		From:        from,
		To:          to,
		Status:      int(req.Status),
		Offset:      int(req.Offset),
		Limit:       int(req.Limit),
	}
}

func Pb2DtoAdminSearchRelatedNewsReq(req *pb.AdminSearchRelatedNewsReq) *dto.AdminSearchRelatedNewsReq {
	if req == nil {
		return nil
	}
	from, _ := time.Parse("2006-01-02 15:04:05", req.From)
	to, _ := time.Parse("2006-01-02 15:04:05", req.To)
	return &dto.AdminSearchRelatedNewsReq{
		Tags:        req.Tags,
		Title:       req.Title,
		Content:     req.Content,
		Category:    int(req.Category),
		SubCategory: int(req.SubCategory),
		PublishBy:   req.PublishBy,
		From:        from,
		To:          to,
		Offset:      int(req.Offset),
		Limit:       int(req.Limit),
	}
}

func Pb2DtoSearchNewsReq(req *pb.SearchNewsReq) *dto.SearchNewsReq {
	if req == nil {
		return nil
	}
	return &dto.SearchNewsReq{
		Keywords: req.Keywords,
		Offset:   int(req.Offset),
		Limit:    int(req.Limit),
	}
}
