package mongodb

import (
	"github.com/farnasirim/shopapi"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// type LineItem struct {
// }

type LineItemBson struct {
	mongodbService *MongodbService   `bson:",skip"`
	IDField        objectid.ObjectID `bson:"_id,omitempty"`
	QuantityField  int               `bson:"quantity"`
	DollarsField   int               `bson:"dollars"`
	CentsField     int               `bson:"cents"`
	ProductIDField objectid.ObjectID `bson:"product_id"`
	OrderIDField   objectid.ObjectID `bson:"order_id"`
}

func (l *LineItemBson) ID() string {
	return l.IDField.Hex()
}

func (l *LineItemBson) Quantity() int {
	return l.QuantityField
}

func (l *LineItemBson) Price() shopapi.DollarValue {
	return newDollarValue(l.DollarsField, l.CentsField)
}

func (l *LineItemBson) ProductID() string {
	return l.ProductIDField.Hex()
}

func (l *LineItemBson) OrderID() string {
	return l.OrderIDField.Hex()
}

func (l *LineItemBson) ProductName() string {
	return l.mongodbService.ProductByID(l.ProductIDField.Hex()).Name()
}
