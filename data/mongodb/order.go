package mongodb

import (
	"github.com/farnasirim/shopapi"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

var (
	orderCollectionName = "orders"

	orderShopIDField = "shop_id"
	orderLinesField  = "lines"
)

type Order struct {
	id     string
	price  *dollarValue
	shopid string
	lines  []*LineItemBson
}

type orderBson struct {
	ID     objectid.ObjectID `bson:"_id,omitempty"`
	ShopID objectid.ObjectID `bson:"shop_id"`
	Lines  []*LineItemBson   `bson:"lines"`
}

func NewOrderFromBson(mongodbService *MongodbService, order *orderBson) *Order {
	return NewOrder(mongodbService, order.ID.Hex(), order.ShopID.Hex(), order.Lines)
}

func NewOrder(mongodbService *MongodbService, id, shopid string, lines []*LineItemBson) *Order {
	return &Order{
		id:     id,
		shopid: shopid,
		lines:  lines,
	}
}

func (o *Order) ID() string {
	return o.id
}

func (o *Order) Lines() []shopapi.LineItem {
	ret := make([]shopapi.LineItem, 0)
	for _, lineItem := range o.lines {
		ret = append(ret, lineItem)
	}
	return ret
}

func (o *Order) Price() shopapi.DollarValue {
	retPrice := newDollarValue(0, 0)

	for _, item := range o.lines {
		retPrice = retPrice.Add(item.Price())
	}
	return retPrice
}

func (o *Order) ShopID() string {
	return o.shopid
}
