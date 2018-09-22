package graphql

import (
	"fmt"

	"github.com/farnasirim/shopapi"
)

func dollarValueModelToGraphQL(dollarValue shopapi.DollarValue) *DollarValue {
	return NewDollarValue(dollarValue.Dollars(), dollarValue.Cents())
}

type DollarValue struct {
	dollars int
	cents   int
}

func NewDollarValue(dollars, cents int) *DollarValue {
	return &DollarValue{
		dollars: dollars,
		cents:   cents,
	}
}

func (d *DollarValue) Display() (string, error) {
	return fmt.Sprintf("$%d.%d", d.dollars, d.cents), nil
}
