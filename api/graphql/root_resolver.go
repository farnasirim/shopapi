package graphql

import (
	"context"
	// graphql "github.com/graph-gophers/graphql-go"
)

type RootResolver struct {
}

func (r *RootResolver) Shops(ctx context.Context) ([]*Shop, error) {
	dataService := dataServiceFromContext(ctx)
	shopsFromDataService := dataService.Shops()

	shopsToReturn := make([]*Shop, 0)

	for _, shop := range shopsFromDataService {
		shopsToReturn = append(shopsToReturn, shopModelToGraphQL(dataService, shop))
	}

	return shopsToReturn, nil
}

type shopParams struct {
	ShopName string
}

func (r *RootResolver) ShopByName(ctx context.Context, params shopParams) (*Shop, error) {
	dataService := dataServiceFromContext(ctx)
	shop := dataService.ShopByName(params.ShopName)
	return shopModelToGraphQL(dataService, shop), nil
}
