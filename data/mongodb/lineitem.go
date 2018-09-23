package mongodb

import (
	"github.com/farnasirim/shopapi"
)

type LineItem struct {
}

func (l *LineItem) ID() string {
	return ""
}

func (l *LineItem) Quantity() int {
	return 0
}

func (l *LineItem) Price() shopapi.DollarValue {
	return nil
}

func (l *LineItem) ProductID() string {
	return ""
}

func (l *LineItem) OrderID() string {
	return ""
}

func (l *LineItem) ProductName() string {
	return ""
}
