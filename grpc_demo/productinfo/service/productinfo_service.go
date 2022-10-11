package main

import (
	"context"
	pb "geekGo/grpc_demo/productinfo/service/ecommerce"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MyServer struct {
	productMap map[string]*pb.Product
}

func (s *MyServer) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {

	out, err := uuid.NewV4()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating", err)
	}

	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	s.productMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()

}

func (s *MyServer) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {

	value, exists := s.productMap[in.Value]

	if exists {
		return value, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Produc dose not exist", in.Value)

}
