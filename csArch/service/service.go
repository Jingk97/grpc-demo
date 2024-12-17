package service

import (
	"context"
)

var ProductService = &productService{}

type productService struct {
}

func (ps *productService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	stock := ps.GetStock(request)
	res := &ProductResponse{ProdStock: stock}
	return res, nil
}

func (ps *productService) GetStock(request *ProductRequest) int32 {
	return request.ProdId
}
