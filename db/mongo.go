package db

import (
	"context"
	"log"
	"reflect"
	"time"

	"github.com/seanburman/game/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient = NewMongoClient(DatabaseOptions{})

// DatabaseOptions stores relevant database names.
type DatabaseOptions struct {
	Db   Databases
	Coll Collections
}

type Databases struct {
	GameDB string
}

type Collections struct {
	UserCollection string
}

type mongoClient struct {
	*mongo.Client
	DatabaseOptions
}

func NewMongoClient(opts DatabaseOptions) *mongoClient {
	md := &mongoClient{DatabaseOptions: opts}
	if err := md.Connect(); err != nil {
		log.Panicln(err.Error())
	}
	return md
}

func (m *mongoClient) Connect() error {
	uri := config.Env().MONGO_URI
	t := reflect.TypeOf(bson.M{})
	reg := bson.NewRegistry()
	reg.RegisterTypeMapEntry(bson.TypeEmbeddedDocument, t)
	// reg := bson.NewRegistryBuilder().RegisterTypeMapEntry(, t).Build()
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(uri).SetTimeout(time.Second*30).SetRegistry(reg),
	)
	if err != nil {
		return err
	}
	m.Client = client
	log.Println("Connected to MongoDB...")
	return nil
}

// Disconnect() should be defered after calling Connect()
func (m *mongoClient) Disconnect() error {
	if err := m.Client.Disconnect(context.TODO()); err != nil {
		return err
	}
	log.Println("Disconnected from MongoDB...")
	return nil
}
