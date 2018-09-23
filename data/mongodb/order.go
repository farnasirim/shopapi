package mongodb

import (
	"github.com/farnasirim/shopapi"
)

type Order struct {
}

func (o *Order) ID() string {
	return ""
}

func (o *Order) Lines() []shopapi.LineItem {
	return nil
}

func (o *Order) Price() shopapi.DollarValue {
	return nil
}

func (o *Order) ShopID() string {
	return ""
}
