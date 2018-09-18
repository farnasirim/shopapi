package graphql

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type createShopParams struct {
	Name string
}

func (r *RootResolver) CreateShop(ctx context.Context, params createShopParams) (*Shop, error) {
	return nil, nil
}

type createProductInShopParams struct {
	ShopID      graphql.ID
	ProductName string
	Dollars     int32
	Cents       int32
}

func (r *RootResolver) CreateProductInShop(ctx context.Context, params createProductInShopParams) (*Product, error) {
	return nil, nil
}

type createOrderInShopParams struct {
	ShopID graphql.ID
}

func (r *RootResolver) CreateOrderInShop(ctx context.Context, params createOrderInShopParams) (*Order, error) {
	return nil, nil
}

type addProductToOrderParams struct {
	ShopID    graphql.ID
	ProductID graphql.ID
	HowMany   int32
}

func (r *RootResolver) AddProductToOrder(ctx context.Context, params addProductToOrderParams) (*LineItem, error) {
	return nil, nil
}
