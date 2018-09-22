package graphql

import (
	_ "context"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/farnasirim/shopapi"
)

func lineItemModelToGraphQL(dataService shopapi.DataService, lineItem shopapi.LineItem) *LineItem {
	return NewLineItem(dataService, lineItem.ID(), lineItem.ProductID(),
		lineItem.Quantity(), dollarValueModelToGraphQL(lineItem.Price()))
}

type LineItem struct {
	id          graphql.ID
	dataService shopapi.DataService
	productID   string
	quantity    int
	price       *DollarValue
}

func NewLineItem(dataService shopapi.DataService, id string, productID string, quantity int, price *DollarValue) *LineItem {
	return &LineItem{
		dataService: dataService,
		id:          graphql.ID(id),
		productID:   productID,
		quantity:    quantity,
		price:       price,
	}
}

func (l *LineItem) Product() (*Product, error) {
	return productModelToGraphQL(l.dataService, l.dataService.ProductByID(l.productID)), nil
}

func (l *LineItem) Quantity() (int32, error) {
	return int32(l.quantity), nil
}

func (l *LineItem) Price() (*DollarValue, error) {
	return l.price, nil
}

func (l *LineItem) ID() (graphql.ID, error) {
	return l.id, nil
}
