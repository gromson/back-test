package message

import (
	"github.com/google/uuid"
	"time"
)

// Message structure representing a message
type Message struct {
	// ID message id
	ID *uuid.UUID `bson:"uuid,omitempty" json:"id"`
	// Name author's name
	Name *string `bson:"name,omitempty" json:"name"`
	// Email author's email
	Email *string `bson:"email,omitempty" json:"email"`
	// Text message text
	Text *string `bson:"text,omitempty" json:"text"`
	// CreationTime message creation time
	CreationTime *time.Time `bson:"creation_time,omitempty" json:"creation_time"`
}
