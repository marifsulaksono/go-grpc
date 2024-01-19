package services

import (
	"context"
	"go-grpc/app/helpers"
	pbPagination "go-grpc/protobuf/pagination"
	pbProduct "go-grpc/protobuf/product"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProductService struct {
	pbProduct.UnimplementedProductServiceServer
	DB *gorm.DB
}

func (p *ProductService) GetProducts(ctx context.Context, pageParam *pbProduct.Page) (*pbProduct.Products, error) {
	var page int64 = 1
	if pageParam.GetPage() != 0 {
		page = pageParam.GetPage()
	}

	var pagination pbPagination.Pagination
	var products []*pbProduct.Product
	sql := p.DB.Table("products AS p").Joins("LEFT JOIN categories AS c ON c.id = p.category_id").
		Select("p.id", "p.name", "p.price", "p.stock", "c.id AS category_id", "c.name")

	offset, limit := helpers.PaginationBuilder(sql, page, &pagination)

	rows, err := sql.Offset(int(offset)).Limit(int(limit)).Rows()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var product pbProduct.Product
		var category pbProduct.Category

		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &category.Id, &category.Name); err != nil {
			log.Fatalf("Failed to get row data : %v", err.Error())
		}

		product.Category = &category
		products = append(products, &product)
	}

	listProducts := &pbProduct.Products{
		Pagination: &pagination,
		Data:       products,
	}

	return listProducts, nil
}
