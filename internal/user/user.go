package user

import "github.com/google/uuid"

type user struct {
	ID       uuid.UUID `bson:"uuid"`
	Name     string    `bson:"name"`
	Password []byte    `bson:"password"`
}

// GetPasswordHash returns clients's hashed password
func (c *user) GetPasswordHash() []byte {
	return c.Password
}

// GetID returns clients's ID
func (c *user) GetID() uuid.UUID {
	return c.ID
}
