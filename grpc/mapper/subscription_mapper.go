package mapper

import (
	"lykafe/news/grpc/news/pb"
	"lykafe/news/core/data/dto"
	"lykafe/news/core/data/model"
	"fmt"
)

func Model2PbSubscription(s *model.Subscription) *pb.Subscription {
	if s == nil {
		return nil
	}
	return &pb.Subscription {
		Email: s.Email,
		Name: s.Name,
		Status: int32(s.Status),
		CreatedAt: s.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func Pb2DtoSearchSubscriptionReq(s *pb.SearchSubscriptionReq) *dto.SearchSubscriptionReq {
	if s == nil {
		return nil
	}
	return &dto.SearchSubscriptionReq {
		Email: s.Email,
		Status: int(s.Status),
		Offset: int(s.Offset),
		Limit: int(s.Limit),
	}
}

func Pb2DtoContactReq(c *pb.ContactReq) *dto.ContactReq {
	if c == nil {
		return nil
	}
	return &dto.ContactReq {
		Email: c.Email,
		Phone: c.Phone,
		Name: c.Name,
		Subject: c.Subject,
		Msg: c.Msg,
		Offset: int(c.Offset),
		Limit: int(c.Limit),
	}
}

func Model2PbContact(c *model.Contact) *pb.Contact {
	if c == nil {
		return nil
	}
	fmt.Println("Model2PbContact ", c)
	return &pb.Contact {
		Id: c.Id,
		Email: c.Email,
		Phone: c.Phone,
		Name: c.Name,
		Subject: c.Subject,
		Msg: c.Msg,
		Status: int32(c.Status),
		CreatedAt: c.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

