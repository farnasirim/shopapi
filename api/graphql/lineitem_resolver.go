package graphql

import (
	_ "context"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/farnasirim/shopapi"
)

func LineItemModelToGraphQL(shopapi.LineItem) *LineItem {
	return nil
}

type LineItem struct {
}

func (l *LineItem) Product() (*Product, error) {
	return nil, nil
}

func (l *LineItem) Quantity() (int32, error) {
	return 0, nil
}

func (l *LineItem) Price() (*DollarValue, error) {
	return nil, nil
}

func (l *LineItem) ID() (graphql.ID, error) {
	return "", nil
}
