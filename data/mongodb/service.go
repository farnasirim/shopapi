package mongodb

import (
	"context"
	"log"

	"github.com/farnasirim/shopapi"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type MongodbService struct {
	db *mongo.Client
}

func NewMongodbService(connectionString string) *MongodbService {
	client, err := mongo.NewClient(connectionString)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = client.Connect(context.TODO())

	mongodbService := &MongodbService{}

	return mongodbService
}

func (s *MongodbService) Shops() []shopapi.Shop {
	return nil
}

func (s *MongodbService) ShopByName(name string) shopapi.Shop {

	return nil
}

func (s *MongodbService) ShopByID(id string) shopapi.Shop {

	return nil
}

func (s *MongodbService) ProductByID(id string) shopapi.Product {

	return nil
}

func (s *MongodbService) ShopOrderByID(shopID, orderID string) shopapi.Order {

	return nil
}

func (s *MongodbService) NewShop(name string) shopapi.Shop {

	return nil
}

func (s *MongodbService) CreateProductInShop(shopID, productName string, dollars, cents int) shopapi.Product {

	return nil
}

func (s *MongodbService) CreateOrderInShop(shopID string) shopapi.Order {
	return nil
}

func (s *MongodbService) AddProductToOrder(orderID, productID string, howMany int) shopapi.LineItem {

	return nil
}
