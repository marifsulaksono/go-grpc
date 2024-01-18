package services

import (
	"context"
	pbPagination "go-grpc/protobuf/pagination"
	pbProduct "go-grpc/protobuf/product"
)

type ProductService struct {
	pbProduct.UnimplementedProductServiceServer
}

func (p *ProductService) GetProducts(context.Context, *pbProduct.Empty) (*pbProduct.Products, error) {
	products := &pbProduct.Products{
		Pagination: &pbPagination.Pagination{
			Total:       50,
			Limit:       10,
			CurrentPage: 1,
			TotalPage:   5,
		},
		Data: []*pbProduct.Product{
			{
				Id:    1,
				Name:  "Laptop Lenovo",
				Price: 5250000.0,
				Stock: 99,
				Category: &pbProduct.Category{
					Id:   1,
					Name: "Laptop & Komputer",
				},
			},
			{
				Id:    2,
				Name:  "Laptop Dell",
				Price: 4850000.0,
				Stock: 90,
				Category: &pbProduct.Category{
					Id:   1,
					Name: "Laptop & Komputer",
				},
			},
		},
	}

	return products, nil
}
