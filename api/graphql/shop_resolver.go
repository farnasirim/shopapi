package graphql

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/farnasirim/shopapi"
)

func shopModelToGraphQL(dataService shopapi.DataService, shop shopapi.Shop) *Shop {
	return NewShop(dataService, shop.ID())
}

type Shop struct {
	dataService shopapi.DataService
	id          graphql.ID
}

func NewShop(dataService shopapi.DataService, id string) *Shop {
	return &Shop{
		dataService: dataService,
		id:          graphql.ID(id),
	}
}

type productsParams struct {
}

func (s *Shop) Products(ctx context.Context) ([]*Product, error) {
	return nil, nil
}

func (s *Shop) Orders(ctx context.Context) ([]*Order, error) {
	return nil, nil
}

func (s *Shop) TotalSales(ctx context.Context) (*DollarValue, error) {
	return nil, nil
}

func (s *Shop) Name() (string, error) {
	return s.Underlaying().Name(), nil
}

func (s *Shop) ID() (graphql.ID, error) {
	return s.id, nil
}

func (s *Shop) Underlaying() shopapi.Shop {
	return s.dataService.ShopByID(string(s.id))
}
