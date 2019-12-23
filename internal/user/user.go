package user

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type user struct {
	ID       primitive.Binary `bson:"uuid"`
	Name     string           `bson:"name"`
	Password []byte           `bson:"password"`
}

// GetPasswordHash returns clients's hashed password
func (c *user) GetPasswordHash() []byte {
	return c.Password
}

// GetID returns clients's ID
func (c *user) GetID() uuid.UUID {
	id, _ := uuid.FromBytes(c.ID.Data)
	return id
}
