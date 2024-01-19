package main

import (
	"log"
	"net"

	"go-grpc/app/config"
	"go-grpc/app/services"
	pbProduct "go-grpc/protobuf/product"

	"google.golang.org/grpc"
)

const port = ":50505"

func main() {
	netListener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error listeing : %v", err)
	}

	grpcServer := grpc.NewServer()
	db := config.Connect()
	productService := services.ProductService{DB: db}
	pbProduct.RegisterProductServiceServer(grpcServer, &productService)
	log.Printf("Server started at %v", netListener.Addr())
	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("Failed to start grpc server : %v", err)
	}
}
