package mongodb

import (
	"github.com/farnasirim/shopapi"
)

type dollarValue struct {
	dollars int
	cents   int
}

func newDollarValue(dollars, cents int) *dollarValue {
	return &dollarValue{
		dollars: dollars,
		cents:   cents,
	}
}

func (d *dollarValue) Dollars() int {
	return d.dollars
}

func (d *dollarValue) Cents() int {
	return d.cents
}

func (d *dollarValue) Add(rhs shopapi.DollarValue) *dollarValue {
	return newDollarValue(
		d.dollars+rhs.Dollars()+(d.cents+rhs.Cents())/100,
		(d.cents+rhs.Cents())%100,
	)
}

func (d *dollarValue) Mul(rhs int) *dollarValue {
	return newDollarValue(
		d.dollars*rhs+(d.cents*rhs)/100,
		(d.cents*rhs)%100,
	)
}
