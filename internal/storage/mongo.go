package storage

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func CreateMongoClient(uri string) (*mongo.Client, error) {
	if uri == "" {
		return nil, errors.New("mongodb connection uri is empty")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		return nil, fmt.Errorf("couldn't create a mongoDB client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		return nil, fmt.Errorf("couldn't connect to MongoDB server: %w", err)
	}

	return client, nil
}
