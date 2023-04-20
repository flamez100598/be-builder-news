package appgrpc

import (
	"lykafe/news/grpc/handler"
	pb "lykafe/news/grpc/news/pb"

	"google.golang.org/grpc"
)

func NewGrpcServer(newsHandler *handler.NewsHandler,
	subscriptionHandler *handler.SubscriptionHandler) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterNewsServiceServer(s, newsHandler)
	pb.RegisterSubscriptionServiceServer(s, subscriptionHandler)
	return s
}
