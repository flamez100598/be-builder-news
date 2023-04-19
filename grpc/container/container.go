package container

import (
	"lykafe/news/grpc/handler"
	"google.golang.org/grpc"
	appgrpc "lykafe/news/grpc"
	"lykafe/news/adapter/mysql"
	"lykafe/news/adapter/kafka"
	"lykafe/news/core/service"
)

func InitializeGrpcServer() *grpc.Server {
	categoryRepo := mysql.NewCategoryRepo()
	newsRepo := mysql.NewNewsRepo()
	newsUsecase := service.NewNewsService(categoryRepo, newsRepo)
	newsHandler := handler.NewNewsHandler(newsUsecase)

	subscriptionRepo := mysql.NewSubscriptionRepo()
	subscriptionUsecase := service.NewSubscriptionService(subscriptionRepo)
	subscriptionHandler := handler.NewSubscriptionHandler(subscriptionUsecase)

	s := appgrpc.NewGrpcServer(newsHandler, subscriptionHandler)
	return s
}

func InitializeKafkaConsumer() {
	mq := kafka.NewKafka()
	userRepo := mysql.NewUserRepo()
	c := service.NewConsumerService(mq, userRepo)
	c.Start()
}
