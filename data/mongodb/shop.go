package mongodb

import (
	"github.com/farnasirim/shopapi"
)

type Shop struct {
}

func (s *Shop) ID() string {

	return ""
}

func (s *Shop) Name() string {

	return ""
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
