package graphql

import (
	_ "context"
	// graphql "github.com/graph-gophers/graphql-go"
)

type Order struct {
}

func (o *Order) Lines() ([]*LineItem, error) {
	return nil, nil
}

func (o *Order) Price() (*DollarValue, error) {
	return nil, nil
}
