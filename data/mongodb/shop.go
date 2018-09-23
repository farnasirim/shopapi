package mongodb

import (
	"github.com/farnasirim/shopapi"
)

var (
	ShopCollectionName = "shops"

	ShopNameField = "name"
)

type Shop struct {
	mongodbService *MongodbService
	id             string
	name           string
}

func NewShop(mongodbService *MongodbService, id, name string) *Shop {
	return &Shop{
		mongodbService: mongodbService,
		id:             id,
		name:           name,
	}
}

func (s *Shop) ID() string {
	return s.id
}

func (s *Shop) Name() string {
	return s.name
}

func (s *Shop) Products() []shopapi.Product {
	return nil
}

func (s *Shop) Orders() []shopapi.Order {
	return nil
}

func (s *Shop) TotalSales() shopapi.DollarValue {
	return nil
}
