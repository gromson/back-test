package server

import (
	"back-api/internal/message"
	"github.com/google/uuid"
)

type Storage interface {
	// Get retrieves a message from a storage by its id
	Get(id uuid.UUID) (*message.Message, error)
	// Fetch fetches all messages from a storage. Latest messages from in the beginning
	Fetch() ([]*message.Message, error)
	// Insert inserts new message into a storage
	Insert(m *message.Message) error
	// Edit modifies the message by given ID. Nil fields will be ignored
	Edit(id uuid.UUID, text string) error
}
