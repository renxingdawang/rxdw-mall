package main

import (
	"context"
	"github.com/renxingdawang/rxdw-mall/server/cmd/product/pkg/mysql"
	"github.com/renxingdawang/rxdw-mall/server/shared/errno"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/product"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct {
	ProductMysqlManager
	CategoryMysqlManager
	ProductRedisManager
}
type ProductMysqlManager interface {
	GetById(ctx context.Context, productId int32) (product mysql.Product, err error)
	SearchProduct(ctx context.Context, q string) (product []*mysql.Product, err error)
}
type CategoryMysqlManager interface {
	GetProductsByCategoryName(ctx context.Context, name string) (category []mysql.Category, err error)
}
type ProductRedisManager interface {
	GetByID(ctx context.Context, productId int32) (product mysql.Product, err error)
}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// TODO: Your code here...
	c, err := s.CategoryMysqlManager.GetProductsByCategoryName(ctx, req.GetCategoryName())
	if err != nil {
		return nil, err
	}
	resp = product.NewListProductsResp()
	for _, v1 := range c {
		for _, v := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{Id: v.ID, Name: v.Name, Description: v.Description, Picture: v.Picture, Price: v.Price})
		}
	}
	return resp, nil
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// TODO: Your code here...
	if req.GetId() == 0 {
		return nil, errno.ProductSrvErr.WithMessage("product id is required")
	}
	p, err := s.ProductRedisManager.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &product.GetProductResp{
		Product: &product.Product{
			Id:          p.ID,
			Picture:     p.Picture,
			Price:       p.Price,
			Description: p.Description,
			Name:        p.Name,
		},
	}, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// TODO: Your code here...
	p, err := s.ProductMysqlManager.SearchProduct(ctx, req.GetQuery())
	var results []*product.Product
	for _, v := range p {
		results = append(results, &product.Product{
			Id:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}
	return &product.SearchProductsResp{Results: results}, err
}
