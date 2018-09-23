package mongodb

import (
	"context"
	"log"

	"github.com/farnasirim/shopapi"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type MongodbService struct {
	db *mongo.Database
}

func NewMongodbService(connectionString, dbName string) *MongodbService {
	client, err := mongo.NewClient(connectionString)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = client.Connect(context.TODO())

	mongodbService := &MongodbService{
		db: client.Database(dbName),
	}

	return mongodbService
}

func (s *MongodbService) EnsureIndices() {
	// TODO: unique on shop name
	// TODO: on shopID of orders
	// TODO: on shopID of products
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
	insertResult, err := s.db.Collection(ShopCollectionName).InsertOne(context.Background(), map[string]string{ShopNameField: name})
	if err != nil {
		// FIXME: handle the error correctly!
		log.Fatalln(err.Error())
	}
	insertedID := insertResult.InsertedID.(objectid.ObjectID).Hex()
	return NewShop(s, insertedID, name)
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
