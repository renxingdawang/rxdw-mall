package main

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/renxingdawang/rxdw-mall/server/cmd/cart/pkg/mysql"
	"github.com/renxingdawang/rxdw-mall/server/shared/errno"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/cart"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/product"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct {
	CartMysqlManager
	ProductManager
}
type CartMysqlManager interface {
	GetCartByUserId(ctx context.Context, userId int32) (cartList []*mysql.Cart, err error)
	AddCart(ctx context.Context, c *mysql.Cart) error
	EmptyCart(ctx context.Context, userId int32) error
}
type ProductManager interface {
	ListProducts(ctx context.Context, req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error)
	GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error)
	SearchProducts(ctx context.Context, req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error)
}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	getProduct, err := s.ProductManager.GetProduct(ctx, &product.GetProductReq{Id: req.Item.GetProductId()})
	if err != nil {
		return nil, err
	}
	if getProduct.Product == nil || getProduct.Product.Id == 0 {
		return nil, errno.CartSrvErr.WithMessage("product not exist")
	}
	err = s.CartMysqlManager.AddCart(ctx, &mysql.Cart{
		UserId:    req.GetUserId(),
		ProductId: req.GetItem().GetProductId(),
		Qty:       req.GetItem().GetQuantity(),
	})
	if err != nil {
		return nil, errno.CartSrvErr.WithMessage("product add error")
	}
	return &cart.AddItemResp{}, nil
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	carts, err := s.CartMysqlManager.GetCartByUserId(ctx, req.GetUserId())
	if err != nil {
		return nil, errno.CartSrvErr.WithMessage("get cart error")
	}
	var items []*cart.CartItem
	for _, v := range carts {
		items = append(items, &cart.CartItem{ProductId: v.ProductId, Quantity: v.Qty})
	}

	return &cart.GetCartResp{Cart: &cart.Cart{UserId: req.GetUserId(), Items: items}}, nil
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	err = s.CartMysqlManager.EmptyCart(ctx, req.GetUserId())
	if err != nil {
		return &cart.EmptyCartResp{}, errno.CartSrvErr.WithMessage("empty cart error")
	}
	return &cart.EmptyCartResp{}, nil
}
