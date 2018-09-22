package graphql

import (
	"context"
	"testing"

	"github.com/farnasirim/shopapi"
	"github.com/farnasirim/shopapi/mock"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestShops(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDataService := mock.NewMockDataService(ctrl)

	shop1 := mock.NewMockShop(ctrl)
	shop1.EXPECT().Name().Return("some-shop")
	shop1.EXPECT().ID().Return("some-id")

	shops := []shopapi.Shop{shop1}

	mockDataService.EXPECT().Shops().Return(shops).Times(1)
	mockDataService.EXPECT().ShopByID("some-id").Return(shop1).Times(1)

	rootResolver := &RootResolver{}

	backgroundContext := context.Background()
	contextWithDataService := context.WithValue(backgroundContext, shopapi.DataServiceStr, mockDataService)

	returnedShops, err := rootResolver.Shops(contextWithDataService)

	assert.Equal(t, 1, len(returnedShops), "return one shop")
	assert.Equal(t, nil, err, "retrieve without error")

	returnedShopName, err := returnedShops[0].Name()
	assert.Equal(t, nil, err, "retrieve the name without error")
	assert.Equal(t, "some-shop", returnedShopName, "with the name of some-shop")
}
