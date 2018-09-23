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
