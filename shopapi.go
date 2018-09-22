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
	ShopID() string
	Price() DollarValue
}

type Order interface {
	ID() string
	Lines() []LineItem
	Price() DollarValue
	ShopID() string
}

type LineItem interface {
	ID() string
	Quantity() int
	Price() DollarValue
	ProductID() string
	ProductName() string
}
type DollarValue interface {
	Dollars() int
	Cents() int
}

type DataService interface {
	Shops() []Shop
	ShopByName(name string) Shop
	ShopByID(id string) Shop

	ProductByID(id string) Product

	ShopOrderByID(shopID, orderID string) Order
}

var (
	DataServiceStr = "data_service"
)
