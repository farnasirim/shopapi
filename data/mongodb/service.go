package mongodb

import (
	"context"
	"log"

	"github.com/farnasirim/shopapi"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	idField string = "_id"
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

func (s *MongodbService) findOneByQuery(collectionName string, query *bson.Document, result interface{}) error {
	err := s.db.Collection(collectionName).FindOne(context.Background(), query).Decode(result)
	return err
}

func (s *MongodbService) ShopByName(name string) shopapi.Shop {
	shop := &shopBson{}
	query := bson.NewDocument(bson.EC.String(shopNameField, name))
	err := s.findOneByQuery(shopCollectionName, query, shop)

	if err != nil {
		// FIXME handle and propagate error correctly
		log.Fatalln(err.Error())
	}

	return NewShopFromBson(s, shop)
}

func (s *MongodbService) ShopByID(id string) shopapi.Shop {
	shop := &shopBson{}
	shopID, err := objectid.FromHex(id)
	if err != nil {
		// FIXME handle and propagate error correctly
		log.Fatalln(err.Error())
	}

	query := bson.NewDocument(bson.EC.ObjectID(idField, shopID))
	err = s.findOneByQuery(shopCollectionName, query, shop)

	if err != nil {
		// FIXME handle and propagate error correctly
		log.Fatalln(err.Error())
	}

	return NewShopFromBson(s, shop)
}

func (s *MongodbService) ProductByID(id string) shopapi.Product {
	productID, err := objectid.FromHex(id)
	if err != nil {
		// FIXME handle and propagate error correctly
		log.Fatalln(err.Error())
	}
	product := &productBson{
		ID: productID,
	}
	query := bson.NewDocument(bson.EC.ObjectID(idField, productID))
	err = s.findOneByQuery(productCollectionName, query, product)

	if err != nil {
		// FIXME handle and propagate error correctly
		log.Fatalln(err.Error())
	}

	return NewProductFromBson(s, product)
}

func (s *MongodbService) OrderByID(shopID, orderID string) shopapi.Order {

	return nil
}

func (s *MongodbService) NewShop(name string) shopapi.Shop {
	shopModel := &shopBson{
		ID:   objectid.New(),
		Name: name,
	}

	insertResult, err := s.db.Collection(shopCollectionName).InsertOne(context.Background(), shopModel)
	if err != nil {
		// FIXME: handle the error correctly!
		log.Fatalln(err.Error())
	}
	shopModel.ID = insertResult.InsertedID.(*bson.Element).Value().ObjectID()
	return NewShopFromBson(s, shopModel)
}

func (s *MongodbService) CreateProductInShop(shopID, productName string, dollars, cents int) shopapi.Product {
	shopIDBson, err := objectid.FromHex(shopID)

	productModel := &productBson{
		// Seems like mongo library won't correctly recognize the empty objectid?
		// And still counts it as if it's present
		// Anyway this is temporary, we need to be able to leave ID empty and
		// it should be the responsibility of the library. I have checked and
		// it intends to do this. Is it my fault? Will have a look later...
		// FIXME
		ID:      objectid.New(),
		Name:    productName,
		ShopID:  shopIDBson,
		Dollars: dollars,
		Cents:   cents,
	}

	insertResult, err := s.db.Collection(productCollectionName).InsertOne(context.Background(), productModel)
	if err != nil {
		// FIXME: handle the error correctly!
		log.Fatalln(err.Error())
	}
	productModel.ID = insertResult.InsertedID.(*bson.Element).Value().ObjectID()

	return NewProductFromBson(s, productModel)
}

func (s *MongodbService) CreateOrderInShop(shopID string) shopapi.Order {
	// shopIDBson, err := objectid.FromHex(shopID)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }
	// orderModel := &orderBson{
	// 	ShopID:  shopIDBson,
	// 	Dollars: 0,
	// 	Cents:   0,
	// }

	// insertResult, err := s.db.Collection(productName).InsertOne(context.Background(), productModel)
	// if err != nil {
	// 	// FIXME: handle the error correctly!
	// 	log.Fatalln(err.Error())
	// }
	// productModel.ID = insertResult.InsertedID.(objectid.ObjectID)

	// return NewProductFromBson(s, productModel)
	return nil
}

func (s *MongodbService) AddProductToOrder(orderID, productID string, howMany int) shopapi.LineItem {

	return nil
}
