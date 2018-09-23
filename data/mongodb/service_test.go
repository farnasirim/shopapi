package mongodb

import (
	"context"
	"testing"

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
