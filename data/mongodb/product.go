package mongodb

import (
	"github.com/farnasirim/shopapi"
)

type Product struct {
}

func (p *Product) ID() string {
	return ""
}

func (p *Product) Name() string {

	return ""
}

func (p *Product) ShopID() string {

	return ""
}

func (p *Product) Price() shopapi.DollarValue {
	return nil
}
