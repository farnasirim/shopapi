package mongodb

import (
	"github.com/farnasirim/shopapi"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

var (
	productCollectionName = "products"

	productNameField    = "name"
	productShopIDField  = "shop_id"
	productDollarsField = "dollars"
	productCentsField   = "cents"
)

type productBson struct {
	ID      objectid.ObjectID `bson:"_id,omitempty"`
	Name    string            `bson:"name"`
	ShopID  objectid.ObjectID `bson:"shop_id"`
	Dollars int               `bson:"dollars"`
	Cents   int               `bson:"cents"`
}

func NewProductFromBson(mongodbService *MongodbService, product *productBson) *Product {
	return NewProduct(mongodbService, product.ID.Hex(), product.Name, product.ShopID.Hex(), newDollarValue(product.Dollars, product.Cents))
}

func NewProduct(mongodbService *MongodbService, id, name, shopid string, price *dollarValue) *Product {
	return &Product{
		mongodbService: mongodbService,
		id:             id,
		name:           name,
		shopid:         shopid,
		price:          price,
	}
}

type Product struct {
	mongodbService *MongodbService
	id             string
	name           string
	shopid         string
	price          *dollarValue
}

func (p *Product) ID() string {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) ShopID() string {
	return p.shopid
}

func (p *Product) Price() shopapi.DollarValue {
	return p.price
}
