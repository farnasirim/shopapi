package graphql

import (
	"context"
	// graphql "github.com/graph-gophers/graphql-go"
)

type Shop struct {
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
	return "", nil
}
