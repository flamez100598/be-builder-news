package appgrpc

import (
	"google.golang.org/grpc"
	pb "lykafe/news/grpc/news/pb"
	"lykafe/news/grpc/handler"
)

func NewGrpcServer(newsHandler *handler.NewsHandler,
	subscriptionHandler *handler.SubscriptionHandler) *grpc.Server{
	s := grpc.NewServer()
	pb.RegisterNewsServiceServer(s, newsHandler)
	pb.RegisterSubscriptionServiceServer(s, subscriptionHandler)
	return s
}