package graphql

import (
	"context"
	graphql "github.com/graph-gophers/graphql-go"
)

type RootResolver struct {
}

func (r *RootResolver) Shops(ctx context.Context) ([]*Shop, error) {
	return nil, nil
}

type shopParams struct {
	ShopID *graphql.ID
}

func (r *RootResolver) Shop(ctx context.Context, params shopParams) (*Shop, error) {
	return nil, nil
}
