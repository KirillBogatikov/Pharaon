package log

import (
	"context"
	config "github.com/Projector-Solutions/Pharaon-config/proxy"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Logger struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewLogger() (*Logger, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Env.MongoURL))
	if err != nil {
		return nil, err
	}

	db := client.Database("proxy")
	col := db.Collection("log")

	return &Logger{
		client:     client,
		database:   db,
		collection: col,
	}, nil
}

func (l *Logger) Log(record *Record) {
	_, err := l.collection.InsertOne(context.Background(), record)
	if err != nil {
		log.Println("log operation failed", err)
	}
}
