package authentication

import (
	"github.com/google/uuid"
)

// AuthenticationData interface returning clients's authentication data
type AuthenticationData interface {
	// GetPasswordHash returns clients's hashed password
	GetPasswordHash() []byte
	// GetID returns clients's ID
	GetID() uuid.UUID
}

// AuthStorage interface returning a authData
type AuthStorage interface {
	// GetByEmail returns an interface that could be casted to AuthenticationData interface
	GetAuthenticationDataByEmail(email string) (AuthenticationData, error)
}

// AuthPasswordService interface for validating a password
type AuthPasswordService interface {
	// Validate validates a given password based on a hash
	Validate(hash, pswd []byte) error
}

// Auth authentication service
type Auth struct {
	storage         AuthStorage
	passwordService AuthPasswordService
}

// NewAuth returns new authentication service
func NewAuth(storage AuthStorage, passwordService AuthPasswordService) *Auth {
	return &Auth{
		storage:         storage,
		passwordService: passwordService,
	}
}

// Authenticate checks if an email and a password are correct and returns clients's ID
func (a *Auth) Authenticate(email, password string) (uuid.UUID, error) {
	data, err := a.storage.GetAuthenticationDataByEmail(email)

	if err != nil {
		return uuid.UUID{}, NewClientNotFoundError(err)
	}

	if err := a.passwordService.Validate(data.GetPasswordHash(), []byte(password)); err != nil {
		return uuid.UUID{}, NewInvalidPasswordError(err)
	}

	return data.GetID(), nil
}
