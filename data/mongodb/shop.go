package mongodb

import (
	"github.com/farnasirim/shopapi"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

var (
	shopCollectionName = "shops"

	shopNameField = "name"
)

type shopBson struct {
	ID   objectid.ObjectID `bson:"_id,omitempty"`
	Name string            `bson:"name"`
}

type Shop struct {
	mongodbService *MongodbService
	id             string
	name           string
}

func NewShopFromBson(mongodbService *MongodbService, shop *shopBson) *Shop {
	return NewShop(mongodbService, shop.ID.Hex(), shop.Name)
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
	shopProducts := s.mongodbService.shopProducts(s.id)
	ret := make([]shopapi.Product, 0)

	for _, prod := range shopProducts {
		ret = append(ret, prod)
	}

	return ret
}

func (s *Shop) Orders() []shopapi.Order {
	ordersOfShop := s.mongodbService.ordersOfShop(s.id)
	ret := make([]shopapi.Order, 0)
	for _, order := range ordersOfShop {
		ret = append(ret, order)
	}

	return ret
}

func (s *Shop) TotalSales() shopapi.DollarValue {
	return s.mongodbService.totalShopSales(s.id)
}
