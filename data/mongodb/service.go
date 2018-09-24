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
	shopIDBson, err := objectid.FromHex(shopID)
	if err != nil {
		log.Fatalln(err.Error())
	}
	orderModel := &orderBson{
		ID:     objectid.New(),
		ShopID: shopIDBson,
		Lines:  make([]*LineItemBson, 0),
	}

	insertResult, err := s.db.Collection(orderCollectionName).InsertOne(context.Background(), orderModel)
	if err != nil {
		// FIXME: handle the error correctly!
		log.Fatalln(err.Error())
	}
	orderModel.ID = insertResult.InsertedID.(*bson.Element).Value().ObjectID()

	return NewOrderFromBson(s, orderModel)
}

func (s *MongodbService) bsonOrderByID(id string) *orderBson {
	orderID, err := objectid.FromHex(id)
	if err != nil {
		// FIXME handle and propagate error correctly
		log.Fatalln(err.Error())
	}
	order := &orderBson{}
	query := bson.NewDocument(bson.EC.ObjectID(idField, orderID))
	err = s.findOneByQuery(orderCollectionName, query, order)

	if err != nil {
		// FIXME handle and propagate error correctly
		log.Fatalln(err.Error())
	}
	return order
}

func (s *MongodbService) OrderByID(id string) shopapi.Order {
	return NewOrderFromBson(s, s.bsonOrderByID(id))
}

func (s *MongodbService) Shops() []shopapi.Shop {
	cur, err := s.db.Collection(shopCollectionName).Find(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	retShops := make([]shopapi.Shop, 0)
	for cur.Next(context.Background()) {
		shop := &shopBson{}
		err := cur.Decode(shop)
		if err != nil {
			log.Fatal(err)
		}
		retShops = append(retShops, NewShopFromBson(s, shop))
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return retShops
}

func (s *MongodbService) AddProductToOrder(orderID, productID string, howMany int) shopapi.LineItem {
	order := s.bsonOrderByID(orderID)

	product := s.ProductByID(productID)
	productObjectID, err := objectid.FromHex(productID)
	if err != nil {
		log.Fatalln(err.Error())
	}

	unitPrice := product.Price()
	price := newDollarValue(unitPrice.Dollars(), unitPrice.Cents()).Mul(howMany)

	// FIXME: use transactions
	lineItem := &LineItemBson{
		IDField:        objectid.New(),
		QuantityField:  howMany,
		DollarsField:   price.Dollars(),
		CentsField:     price.Cents(),
		ProductIDField: productObjectID,
		OrderIDField:   order.ID,
	}
	order.Lines = append(order.Lines, lineItem)

	_, err = s.db.Collection(orderCollectionName).UpdateOne(context.Background(),
		bson.NewDocument(bson.EC.ObjectID(idField, order.ID)),
		bson.NewDocument(bson.EC.SubDocument("$set",
			bson.NewDocument(bson.EC.Interface(orderLinesField, order.Lines)))))
	if err != nil {
		log.Fatalln(err.Error())
	}
	return lineItem
}

func (s *MongodbService) ordersOfShop(shopID string) []*Order {
	shopObjectID, err := objectid.FromHex(shopID)
	if err != nil {
		// FIXME: handle and propagate the error correctly
		log.Fatalln(err.Error())
	}
	cur, err := s.db.Collection(orderCollectionName).Find(context.Background(),
		bson.NewDocument(bson.EC.ObjectID(orderShopIDField, shopObjectID)))
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	retOrders := make([]*Order, 0)
	for cur.Next(context.Background()) {
		order := &orderBson{}
		err := cur.Decode(order)
		if err != nil {
			log.Fatal(err)
		}
		retOrders = append(retOrders, NewOrderFromBson(s, order))
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return retOrders
}

func (s *MongodbService) totalShopSales(shopID string) *dollarValue {
	ret := newDollarValue(0, 0)
	for _, order := range s.ordersOfShop(shopID) {
		ret = ret.Add(order.Price())
	}
	return ret
}

func (s *MongodbService) shopProducts(shopID string) []*Product {
	shopObjectID, err := objectid.FromHex(shopID)
	if err != nil {
		// FIXME: handle and propagate the error correctly
		log.Fatalln(err.Error())
	}
	cur, err := s.db.Collection(productCollectionName).Find(context.Background(),
		bson.NewDocument(bson.EC.ObjectID(orderShopIDField, shopObjectID)))
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	retProds := make([]*Product, 0)
	for cur.Next(context.Background()) {
		prod := &productBson{}
		err := cur.Decode(prod)
		if err != nil {
			log.Fatal(err)
		}
		retProds = append(retProds, NewProductFromBson(s, prod))
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return retProds
}
