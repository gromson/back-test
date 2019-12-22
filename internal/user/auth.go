package user

import (
	"back-api/internal/authentication"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	defaultTimeout  = 2 * time.Second
	usersCollection = "users"
)

// Storage struct for
type Storage struct {
	client *mongo.Client
	dbName string
}

func NewStorage(client *mongo.Client, dbName string) (*Storage, error) {
	if client == nil {
		return nil, errors.New("mongodb client hasn't been provided")
	}

	if dbName == "" {
		return nil, errors.New("db name hasn't been provided")
	}

	return &Storage{
		client: client,
		dbName: dbName,
	}, nil
}

// GetAuthenticationDataByEmail returns an interface that could be casted to AuthenticationData interface
func (s *Storage) GetAuthenticationDataByEmail(email string) (authentication.AuthenticationData, error) {
	usr := new(user)

	c := s.client.Database(s.dbName).Collection(usersCollection)

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	err := c.FindOne(ctx, bson.M{"name": email}).Decode(usr)

	if err == mongo.ErrNoDocuments {
		return usr, fmt.Errorf("user \"%s\" doesn't exists", email)
	}

	if err != nil {
		return usr, fmt.Errorf("error while getting a user: %w", err)
	}

	return usr, nil
}
