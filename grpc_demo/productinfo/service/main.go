package main

import (
	"context"
	pb "geekGo/grpc_demo/productinfo/service/ecommerce"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {

	listener, err := net.Listen("tcp", port)

	if err != nil {

		log.Fatalf("error listening", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(HelloInterceptor),
	}

	s := grpc.NewServer(opts...)

	pb.RegisterProductInfoServer(s, &MyServer{})

	log.Printf("start prot" + port)

	if err := s.Serve(listener); err != nil {

		log.Fatalf("server failed error: %v", err)
	}

}

func HelloInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	log.Println("你好")
	resp, err := handler(ctx, req)
	log.Println("再见")

	return resp, err

}
