package graphql

import (
	_ "context"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/farnasirim/shopapi"
)

func OrderModelToGraphQL(shopapi.Order) *Order {
	return nil
}

type Order struct {
}

func (o *Order) Lines() ([]*LineItem, error) {
	return nil, nil
}

func (o *Order) Price() (*DollarValue, error) {
	return nil, nil
}

func (o *Order) ID() (graphql.ID, error) {
	return "", nil
}
