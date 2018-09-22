package graphql

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/farnasirim/shopapi"
)

func productModelToGraphQL(shopapi.Product) *Product {
	return nil
}

type Product struct {
}

func (p *Product) LinesInOrders(ctx context.Context) ([]*LineItem, error) {
	return nil, nil
}

func (p *Product) Price() (*DollarValue, error) {
	return nil, nil
}

func (p *Product) TotalSold() (*DollarValue, error) {
	return nil, nil
}

func (p *Product) Name() (string, error) {
	return "", nil
}

func (p *Product) ID() (graphql.ID, error) {
	return "", nil
}
