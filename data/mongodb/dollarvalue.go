package mongodb

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
