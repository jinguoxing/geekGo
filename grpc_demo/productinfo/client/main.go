package main

import (
	"context"
	"log"
	"time"

	pb "geekGo/grpc_demo/productinfo/client/ecommerce"

	"google.golang.org/grpc"
)

const (
	addrees = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(addrees, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("connect %v", err)
	}

	defer conn.Close()

	c := pb.NewProductInfoClient(conn)

	name := "Apple 13"
	description := "apple 13 description"
	price := float32(1000.0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // cance
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("add product %v: %v", name, err)
	}

	log.Printf("Product ID %v", r.Value)

	product, _ := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})

	log.Printf("Product  %v", product.String())

} 
