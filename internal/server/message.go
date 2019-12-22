package server

import (
	"back-api/internal/message"
	"errors"
	"github.com/google/uuid"
	"reflect"
)

// ErrEmptyMessageID error stands that message doesn't contain an ID
var ErrEmptyMessageID = errors.New("message should contain an ID")

// ErrEmptyMessageName error stands that message doesn't contain a name
var ErrEmptyMessageName = errors.New("message should contain a name of the author")

// ErrEmptyMessageEmail error stands that message doesn't contain an email
var ErrEmptyMessageEmail = errors.New("message should contain an email of the author")

// ErrEmptyMessageText error stands that message doesn't contain a text
var ErrEmptyMessageText = errors.New("message should contain a text")

// ErrEmptyMessageTime error stands that message doesn't contain a creation time
var ErrEmptyMessageTime = errors.New("message should contain a creation time")

func validateMessage(msg *message.Message) []error {
	errs := make([]error, 0)

	if msg.ID == nil || reflect.DeepEqual(*msg.ID, uuid.UUID{}){
		errs = append(errs, ErrEmptyMessageID)
	}

	if msg.Name == nil || *msg.Name == "" {
		errs = append(errs, ErrEmptyMessageName)
	}

	if msg.Email == nil || *msg.Email == "" {
		errs = append(errs, ErrEmptyMessageEmail)
	}

	if msg.Text == nil || *msg.Text == "" {
		errs = append(errs, ErrEmptyMessageText)
	}

	if msg.CreationTime == nil {
		errs = append(errs, ErrEmptyMessageTime)
	}

	return errs
}
