package message

import (
	"github.com/google/uuid"
	"time"
)

// Message structure representing a message
type Message struct {
	// ID message id
	ID *uuid.UUID `bson:"uuid,omitempty"`
	// Name author's name
	Name *string `bson:"name,omitempty"`
	// Email author's email
	Email *string `bson:"email,omitempty"`
	// Text message text
	Text *string `bson:"text,omitempty"`
	// CreationTime message creation time
	CreationTime *time.Time `bson:"creation_time,omitempty"`
}
