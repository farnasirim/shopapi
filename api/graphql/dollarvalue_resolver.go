package graphql

import (
	_ "context"
	// graphql "github.com/graph-gophers/graphql-go"
)

type DollarValue struct {
}

func (d *DollarValue) Display() (string, error) {
	return "", nil
}
