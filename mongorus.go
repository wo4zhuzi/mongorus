package mongorus

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

type M bson.M

type MongoHook struct {
	MongoCollection *mongo.Collection
}

func NewMongoHook(mongoUrl, db, collection string) (*MongoHook, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s",mongoUrl))

	var err error
	MongoClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return &MongoHook{MongoCollection: MongoClient.Database(db).Collection(collection)}, nil
}

func NewAuthMongoHook(mongoUrl, db, collection string, auth options.Credential) (*MongoHook, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s",mongoUrl)).SetAuth(auth)

	var err error
	MongoClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return &MongoHook{MongoCollection: MongoClient.Database(db).Collection(collection)}, nil
}

func (h *MongoHook)Fire(entry *logrus.Entry) error{
	data := make(logrus.Fields)
	data["Level"] = entry.Level.String()
	data["Time"] = entry.Time.Format("2006-01-02 15:04:05")
	data["Message"] = entry.Message

	for k, filed := range entry.Data {
		data[k] = filed
	}

	_, err := h.MongoCollection.InsertOne(context.Background(), M(data))

	if err != nil {
		return err
	}
	return nil
}

func (h *MongoHook) Levels() []logrus.Level {
	return logrus.AllLevels
}