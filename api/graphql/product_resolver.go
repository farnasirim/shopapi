package graphql

import (
	graphql "github.com/graph-gophers/graphql-go"

	"github.com/farnasirim/shopapi"
)

func productModelToGraphQL(dataService shopapi.DataService, product shopapi.Product) *Product {
	return nil
}

type Product struct {
	dataService shopapi.DataService
	id          graphql.ID
	name        string
	shopID      string
	price       *DollarValue
}

func NewProduct(dataService shopapi.DataService, id, name, shopID string, price *DollarValue) *Product {
	return &Product{
		dataService: dataService,
		id:          graphql.ID(id),
		name:        name,
		shopID:      shopID,
		price:       price,
	}
}

func (p *Product) Price() (*DollarValue, error) {
	return p.price, nil
}

func (p *Product) Name() (string, error) {
	return p.name, nil
}

func (p *Product) ID() (graphql.ID, error) {
	return p.id, nil
}
