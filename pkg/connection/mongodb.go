package connection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(dsn string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func PingMongoDB(client *mongo.Client) error {
	return client.Ping(context.Background(), nil)
}
