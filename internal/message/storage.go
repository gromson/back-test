package message

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	defaultTimeout    = 2 * time.Second
	messageCollection = "messages"
)

// Storage service for working with messages
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

// Get retrieves a message from a storage by its id
func (s *Storage) Get(id uuid.UUID) (*Message, error) {
	msg := &Message{}

	c := s.client.Database(s.dbName).Collection(messageCollection)

	ctx, cancel := newCtxWithTimeout(defaultTimeout)
	defer cancel()

	err := c.FindOne(ctx, bson.M{"uuid": id}).Decode(msg)

	if err == mongo.ErrNoDocuments {
		return msg, fmt.Errorf("message \"%s\" doesn't exists", id)
	}

	if err != nil {
		return msg, fmt.Errorf("error while getting a message: %w", err)
	}

	return msg, nil
}

// Fetch fetches all messages from a storage. Latest messages from in the beginning
func (s *Storage) Fetch() ([]*Message, error) {
	msgs := make([]*Message, 0, 100)

	c := s.client.Database(s.dbName).Collection(messageCollection)

	ctx, cancel := newCtxWithTimeout(defaultTimeout)
	defer cancel()

	cur, err := c.Find(ctx, bson.M{}, options.Find().SetSort(bson.M{"creation_time": -1}))

	if err != nil && err != mongo.ErrNoDocuments {
		return msgs, fmt.Errorf("error while trying to fetch all messages: %w", err)
	}

	if cur == nil {
		return msgs, errors.New("mongodb cursor is nil")
	}

	for cur.Next(context.Background()) {
		msg := new(Message)

		if err := cur.Decode(msg); err != nil {
			return msgs, fmt.Errorf("error while decoding messgae: %w", err)
		}

		msgs = append(msgs, msg)
	}

	if err := cur.Err(); err != nil {
		log.Print("mongodb cursor error:" + err.Error())
	}

	if err := cur.Close(context.Background()); err != nil {
		log.Print("error while closing the mongodb cursor:" + err.Error())
	}

	return msgs, nil
}

// Insert inserts new message into a storage
func (s *Storage) Insert(m *Message) error {
	c := s.client.Database(s.dbName).Collection(messageCollection)

	ctx, cancel := newCtxWithTimeout(defaultTimeout)
	defer cancel()

	_, err := c.InsertOne(ctx, m)

	if err != nil {
		return fmt.Errorf("error while adding a message: %w", err)
	}

	return nil
}

// Edit modifies the message by given ID. Nil fields will be ignored
func (s *Storage) Edit(id uuid.UUID, text string) error {
	c := s.client.Database(s.dbName).Collection(messageCollection)

	ctx, cancel := newCtxWithTimeout(defaultTimeout)
	defer cancel()

	_, err := c.UpdateOne(ctx, bson.M{"uuid": id}, bson.M{"text": text})

	if err != nil {
		return fmt.Errorf("error while adding a message: %w", err)
	}

	return nil
}

func newCtxWithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}
