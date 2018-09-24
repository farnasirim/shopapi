package graphql

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type createShopParams struct {
	Name string
}

func (r *RootResolver) CreateShop(ctx context.Context, params createShopParams) (*Shop, error) {
	dataService := dataServiceFromContext(ctx)
	shopModel := dataService.NewShop(params.Name)
	return shopModelToGraphQL(dataService, shopModel), nil
}

type createProductInShopParams struct {
	ShopID      graphql.ID
	ProductName string
	Dollars     int32
	Cents       int32
}

func (r *RootResolver) CreateProductInShop(ctx context.Context, params createProductInShopParams) (*Product, error) {
	dataService := dataServiceFromContext(ctx)
	productModel := dataService.CreateProductInShop(string(params.ShopID), params.ProductName, int(params.Dollars), int(params.Cents))
	graphqlProduct := productModelToGraphQL(dataService, productModel)
	return graphqlProduct, nil
}

type createOrderInShopParams struct {
	ShopID graphql.ID
}

func (r *RootResolver) CreateOrderInShop(ctx context.Context, params createOrderInShopParams) (*Order, error) {
	dataService := dataServiceFromContext(ctx)
	orderModel := dataService.CreateOrderInShop(string(params.ShopID))
	return orderModelToGraphQL(dataService, orderModel), nil
}

type addProductToOrderParams struct {
	OrderID   graphql.ID
	ProductID graphql.ID
	HowMany   int32
}

func (r *RootResolver) AddProductToOrder(ctx context.Context, params addProductToOrderParams) (*LineItem, error) {
	dataService := dataServiceFromContext(ctx)
	lineItemModel := dataService.AddProductToOrder(string(params.OrderID), string(params.ProductID), int(params.HowMany))
	return lineItemModelToGraphQL(dataService, lineItemModel), nil
}
