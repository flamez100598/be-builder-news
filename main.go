package main

import (
	"log"
	"net"
	"lykafe/news/grpc/container"
	"lykafe/news/config"
)

func main() {
	port := config.ServerPort()
	log.Printf("Starting news server at port " + port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	container.InitializeKafkaConsumer()
	s := container.InitializeGrpcServer()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Printf("Started news server at port " + port)
}