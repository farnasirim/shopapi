package mongodb

import (
	"context"
	"testing"

	"github.com/farnasirim/shopapi"

	"github.com/stretchr/testify/assert"
)

func TestNewShop(t *testing.T) {
	mongodbService := NewMongodbService("mongodb://localhost:27017", "testtest")
	defer mongodbService.db.Drop(context.Background())

	shop := mongodbService.NewShop("some-shop")
	assert.Equal(t, "some-shop", shop.Name())
}

func TestWriteReadShop(t *testing.T) {
	mongodbService := NewMongodbService("mongodb://localhost:27017", "testtest")
	defer mongodbService.db.Drop(context.Background())

	createdShop := mongodbService.NewShop("some-shop")
	foundShop := mongodbService.ShopByName("some-shop")

	assert.Equal(t, createdShop.ID(), foundShop.ID())
	assert.Equal(t, createdShop.Name(), foundShop.Name())
}

func TestCreateProductInShop(t *testing.T) {
	mongodbService := NewMongodbService("mongodb://localhost:27017", "testtest")
	defer mongodbService.db.Drop(context.Background())

	_ = mongodbService.NewShop("some-shop")
	foundShop := mongodbService.ShopByName("some-shop")

	createdProduct := mongodbService.CreateProductInShop(foundShop.ID(), "some-product", 10, 20)

	returnedProduct := mongodbService.ProductByID(createdProduct.ID())

	assert.Equal(t, createdProduct.ID(), returnedProduct.ID())
	assert.Equal(t, createdProduct.Name(), returnedProduct.Name())
	assert.Equal(t, createdProduct.ShopID(), returnedProduct.ShopID())
	assert.Equal(t, createdProduct.Price(), returnedProduct.Price())
}

func TestCreateOrderInShop(t *testing.T) {
	mongodbService := NewMongodbService("mongodb://localhost:27017", "testtest2")
	defer mongodbService.db.Drop(context.Background())

	_ = mongodbService.NewShop("some-shop")
	s := mongodbService.ShopByName("some-shop")

	p1 := mongodbService.CreateProductInShop(s.ID(), "some-product", 10, 20)
	p2 := mongodbService.CreateProductInShop(s.ID(), "other-product", 100, 1)

	order := mongodbService.CreateOrderInShop(s.ID())
	o := mongodbService.OrderByID(order.ID())
	assert.Equal(t, order.ID(), o.ID())

	l1 := mongodbService.AddProductToOrder(o.ID(), p1.ID(), 2)
	l2 := mongodbService.AddProductToOrder(o.ID(), p1.ID(), 1)
	l3 := mongodbService.AddProductToOrder(o.ID(), p2.ID(), 3)

	assert.Equal(t, o.ID(), l1.OrderID())
	assert.Equal(t, o.ID(), l2.OrderID())
	assert.Equal(t, o.ID(), l3.OrderID())

	assert.Equal(t, newDollarValue(330, 63), s.TotalSales())
	assert.Equal(t, newDollarValue(330, 63), mongodbService.OrderByID(order.ID()).Price())
	shopOrders := s.Orders()
	assert.Equal(t, 1, len(shopOrders))
}

func Test_ordersOfShop(t *testing.T) {
	mongodbService := NewMongodbService("mongodb://localhost:27017", "testtest")
	defer mongodbService.db.Drop(context.Background())

	_ = mongodbService.NewShop("some-shop")
	s := mongodbService.ShopByName("some-shop")

	mongodbService.CreateProductInShop(s.ID(), "some-product", 10, 20)
	mongodbService.CreateProductInShop(s.ID(), "other-product", 100, 1)

	order := mongodbService.CreateOrderInShop(s.ID())
	o := mongodbService.OrderByID(order.ID())

	order2 := mongodbService.CreateOrderInShop(s.ID())
	o2 := mongodbService.OrderByID(order2.ID())

	orders := s.Orders()
	assert.Equal(t, 2, len(orders))
	oneExists := orders[0].ID() == o.ID() || orders[1].ID() == o.ID()
	twoExists := orders[0].ID() == o2.ID() || orders[1].ID() == o2.ID()

	assert.True(t, oneExists)
	assert.True(t, twoExists)
}

func Test_dataservie(t *testing.T) {
	var _ shopapi.DataService = NewMongodbService("mongodb://localhost:27017", "testtest")
}

func Test_shopProducts(t *testing.T) {
	mongodbService := NewMongodbService("mongodb://localhost:27017", "testtest")
	defer mongodbService.db.Drop(context.Background())

	_ = mongodbService.NewShop("some-shop")
	s := mongodbService.ShopByName("some-shop")

	p1 := mongodbService.CreateProductInShop(s.ID(), "some-product", 10, 20)
	p2 := mongodbService.CreateProductInShop(s.ID(), "other-product", 100, 1)

	prods := s.Products()
	assert.Equal(t, 2, len(prods))

	oneExists := prods[0].ID() == p1.ID() || prods[1].ID() == p1.ID()
	twoExists := prods[0].ID() == p2.ID() || prods[1].ID() == p2.ID()

	assert.True(t, oneExists)
	assert.True(t, twoExists)
}

func TestShops(t *testing.T) {
	mongodbService := NewMongodbService("mongodb://localhost:27017", "testtest")
	defer mongodbService.db.Drop(context.Background())

	shops := mongodbService.Shops()
	assert.Equal(t, 0, len(shops))

	_ = mongodbService.NewShop("some-shop")
	s := mongodbService.ShopByName("some-shop")

	_ = mongodbService.NewShop("other-shop")
	s2 := mongodbService.ShopByName("other-shop")

	shops = mongodbService.Shops()
	assert.Equal(t, 2, len(shops))

	oneExists := shops[0].ID() == s.ID() || shops[1].ID() == s.ID()
	twoExists := shops[0].ID() == s2.ID() || shops[1].ID() == s2.ID()

	assert.True(t, oneExists)
	assert.True(t, twoExists)
}
