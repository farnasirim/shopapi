//go:generate mockgen --source=shopapi.go --destination=mock/mock.go --package=mock

package shopapi

type Shop interface {
	ID() string
	Name() string
	Products() []Product
	Orders() []Order
	TotalSales() DollarValue
}

type Product interface {
	ID() string
	Name() string
	LinesInOrders() []LineItem
	Price() DollarValue
	TotalSold() DollarValue
}

type Order interface {
	ID() string
	Lines() []LineItem
	Price() DollarValue
}

type LineItem interface {
	ID() string
	Quantity() int
	Price() DollarValue
}
type DollarValue interface {
	Dollars() int
	Cents() int
}

type DataService interface {
	Shops() []Shop
	ShopByName(name string) Shop
	ShopByID(id string) Shop
}

var (
	DataServiceStr = "data_service"
)
