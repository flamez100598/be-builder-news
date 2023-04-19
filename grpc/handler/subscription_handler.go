package handler

import (
	"context"
	pb "lykafe/news/grpc/news/pb"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"lykafe/news/core/port/usecase"
	"lykafe/news/grpc/mapper"
	"lykafe/news/core/data/codes"
	// "fmt"
)

type SubscriptionHandler struct {
	pb.UnimplementedSubscriptionServiceServer
	subscriptionUsecase usecase.Subscription
}

func NewSubscriptionHandler(subscriptionUsecase usecase.Subscription) *SubscriptionHandler{
	return &SubscriptionHandler{
		subscriptionUsecase: subscriptionUsecase,
	}
}

func (h *SubscriptionHandler) Subscribe(ctx context.Context, in *pb.NewsEmail) (*pb.SubscriptionResp, error) {
	sub, code, err := h.subscriptionUsecase.Subscribe(in.Email)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.SubscriptionResp{
		Code:  code.ToInt32(),
		Msg: code.Msg(),
		Subscription: mapper.Model2PbSubscription(sub),
	}, nil
}

func (h *SubscriptionHandler) SearchSubscription(ctx context.Context, in *pb.SearchSubscriptionReq) (*pb.SubscriptionsResp, error) {
	req := mapper.Pb2DtoSearchSubscriptionReq(in)
	subs, err := h.subscriptionUsecase.SearchSubscription(req)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbSubs:= []*pb.Subscription{}
	for _, sub := range subs {
		pbSubs = append(pbSubs, mapper.Model2PbSubscription(sub))
	}
	return  &pb.SubscriptionsResp{
		Code:  codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Subscriptions: pbSubs,
	}, nil
}

func (h *SubscriptionHandler) SearchSubscriptionCount(ctx context.Context, in *pb.SearchSubscriptionReq) (*pb.NewsCountResp, error) {
	c, err := h.subscriptionUsecase.SearchSubscriptionCount(mapper.Pb2DtoSearchSubscriptionReq(in))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.NewsCountResp{
		Code:  codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Count: int32(c),
	}, nil
}

func (h *SubscriptionHandler) DeleteSubscription (ctx context.Context, in *pb.NewsEmail) (*pb.SubscriptionResp, error) {
	s, err := h.subscriptionUsecase.DeleteSubscription(in.Email)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.SubscriptionResp{
		Code:  codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Subscription: &pb.Subscription{
			Email: s.Email,
		},
	}, nil
}

func (h *SubscriptionHandler) SendContact(ctx context.Context, in *pb.ContactReq) (*pb.ContactResp, error) {
	con, err := h.subscriptionUsecase.SendContact(mapper.Pb2DtoContactReq(in))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.ContactResp{
		Code:  codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Contact: mapper.Model2PbContact(con),
	}, nil
}

func (h *SubscriptionHandler) SearchContact(ctx context.Context, in *pb.ContactReq) (*pb.ContactsResp, error) {
	req := mapper.Pb2DtoContactReq(in)
	contacts, err := h.subscriptionUsecase.SearchContact(req)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	pbContacts:= []*pb.Contact{}
	for _, con := range contacts {
		pbContacts = append(pbContacts, mapper.Model2PbContact(con))
	}
	return  &pb.ContactsResp{
		Code:  codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Contacts: pbContacts,
	}, nil
}

func (h *SubscriptionHandler) SearchContactCount(ctx context.Context, in *pb.ContactReq) (*pb.NewsCountResp, error) {
	c, err := h.subscriptionUsecase.SearchContactCount(mapper.Pb2DtoContactReq(in))
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.NewsCountResp{
		Code:  codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		Count: int32(c),
	}, nil
}

func (h *SubscriptionHandler) GetPageContent(ctx context.Context, in *pb.PageContentReq) (*pb.PageContentResp, error) {
	page, err := h.subscriptionUsecase.GetPageContent(in.Page)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.PageContentResp{
		Code:  codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		PageContent: &pb.PageContent{
			Id: int32(page.Id),
			Title: page.Title,
			Content: page.Content,
			UpdatedAt: page.UpdatedAt.Format("2006-01-02 15:04:05"),
			UsernameUpdatedBy: page.UsernameUpdatedBy,
			NameUpdatedBy: page.NameUpdatedBy,
			EmailUpdatedBy: page.EmailUpdatedBy,
			AvatarUpdatedBy: page.AvatarUpdatedBy,
		},
	}, nil
}

func (h *SubscriptionHandler) UpdatePageContent (ctx context.Context, in *pb.PageContentReq) (*pb.PageContentResp, error) {
	page, err := h.subscriptionUsecase.UpdatePageContent(in.Content, in.UpdatedBy, in.Page)
	if err != nil {
		return nil, status.Error(grpcCodes.Internal, err.Error())
	}
	return  &pb.PageContentResp{
		Code:  codes.Ok.ToInt32(),
		Msg: codes.Ok.Msg(),
		PageContent: &pb.PageContent{
			Id: int32(page.Id),
			Title: page.Title,
			Content: page.Content,
			UpdatedAt: page.UpdatedAt.Format("2006-01-02 15:04:05"),
			UsernameUpdatedBy: page.UsernameUpdatedBy,
			NameUpdatedBy: page.NameUpdatedBy,
			EmailUpdatedBy: page.EmailUpdatedBy,
			AvatarUpdatedBy: page.AvatarUpdatedBy,
		},
	}, nil
}
