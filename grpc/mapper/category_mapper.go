package mapper

import (
	"lykafe/news/core/data/model"
	"lykafe/news/grpc/news/pb"
)

func Model2PbCategory(cat *model.Category) *pb.Category {
	if cat == nil {
		return nil
	}
	return &pb.Category {
		Id: int32(cat.Id),
		ParentId: int32(cat.ParentId),
		Name: cat.Name,
		Description: cat.Description,
		Slug: cat.Slug,
		NewsCount: int32(cat.NewsCount),
		Status: int32(cat.Status),
		Order: int32(cat.Order),
	}
}

func Pb2ModelCategory(cat *pb.Category) *model.Category {
	if cat == nil {
		return nil
	}
	return &model.Category {
		Id: int(cat.Id),
		ParentId: int(cat.ParentId),
		Name: cat.Name,
		Description: cat.Description,
		Slug: cat.Slug,
		NewsCount: int(cat.NewsCount),
		Status: int(cat.Status),
		Order: int(cat.Order),
	}
}