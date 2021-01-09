package tools

import (
	"context"
	"fmt"
	"time"

	config "dev/interview-craft/configs"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient interface to define functions
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . MongoClient
type MongoClient interface {
	AddProductToInventory(db, collection string, productData interface{}) (*mongo.InsertOneResult, error)
	DisplayInventory(db, collection string, res interface{}) error
	UpdateInventory()
}

// MongoClientHandler is implementation of MongoClient interface
type MongoClientHandler struct {
	Client  *mongo.Client
	Context context.Context
}

// NewMongoClient
func NewMongoClient(cfg config.Config) MongoClient {
	log.Info("mongo-client-new-mongo-client")
	defer log.Info("mongo-client-new-mongo-client-complete")

	url := fmt.Sprintf("%s:%s", cfg.DB.Server, cfg.DB.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Error("mongo-client-new-client-error")
		return nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Error("mongo-client-connect-error")
		return nil
	}

	return &MongoClientHandler{
		Client:  client,
		Context: context.TODO(),
	}
}

//SetCollection returns the desired MongoDB collection
func (m *MongoClientHandler) SetCollection(db, collection string) *mongo.Collection {
	return m.Client.Database(db).Collection(collection)
}

// AddProductToInventory
func (m *MongoClientHandler) AddProductToInventory(db, collection string, productData interface{}) (*mongo.InsertOneResult, error) {
	log.Info("mongo-client-add-product-to-inventory")
	defer log.Info("mongo-client-add-product-to-inventory-complete")

	// Set the collection
	coll := m.SetCollection(db, collection)
	res, err := coll.InsertOne(m.Context, productData)
	if err != nil {
		log.Error("mongo-client-add-product-to-inventory-insert-one-error ", err)
		return res, err
	}

	return res, err
}

// DisplayInventory
func (m *MongoClientHandler) DisplayInventory(db, collection string, res interface{}) error {
	log.Info("mongo-client-display-inventory")
	defer log.Info("mongo-client-display-inventory-complete")

	coll := m.SetCollection(db, collection)
	cur, err := coll.Find(m.Context, bson.D{})
	if err != nil {
		log.Error("mongo-client-display-inventory-error ", err)
		return err
	}

	err = cur.All(m.Context, res)
	if err != nil {
		log.Error("mongo-client-display-inventory-cursor-all-error ", err)
		return err
	}

	return nil
}

// UpdateInventory
func (m *MongoClientHandler) UpdateInventory() {}
