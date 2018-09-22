package graphql

import (
	_ "context"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/farnasirim/shopapi"
)

func orderModelToGraphQL(dataService shopapi.DataService, order shopapi.Order) *Order {

	// Should we move the list conversions into corresponding *_resolver.go files?
	linesModels := order.Lines()
	graphqlLines := make([]*LineItem, 0)
	for _, lineModel := range linesModels {
		graphqlLines = append(graphqlLines, lineItemModelToGraphQL(dataService, lineModel))
	}
	graphqlDollarValue := dollarValueModelToGraphQL(order.Price())
	return NewOrder(dataService, order.ID(), graphqlLines, graphqlDollarValue, order.ShopID())
}

type Order struct {
	id          graphql.ID
	dataService shopapi.DataService
	lines       []*LineItem
	price       *DollarValue
	shopID      string
}

func NewOrder(dataService shopapi.DataService, id string, lines []*LineItem, price *DollarValue, shopID string) *Order {
	return &Order{
		dataService: dataService,
		id:          graphql.ID(id),
		lines:       lines,
		price:       price,
		shopID:      shopID,
	}
}

func (o *Order) Lines() ([]*LineItem, error) {
	return o.lines, nil
}

func (o *Order) Price() (*DollarValue, error) {
	return o.price, nil
}

func (o *Order) ID() (graphql.ID, error) {
	return o.id, nil
}
