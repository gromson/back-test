package authentication

import (
	"errors"
	"github.com/google/uuid"
)

type authData struct {
	id       uuid.UUID
	password []byte
}

func (p *authData) GetPasswordHash() []byte {
	return p.password
}

func (p *authData) GetID() uuid.UUID {
	return p.id
}

type testAuthStorage map[string]*authData

func (s testAuthStorage) GetAuthenticationDataByEmail(email string) (AuthenticationData, error) {
	p, ok := s[email]

	if !ok {
		return nil, errors.New("wrong email given")
	}

	return p, nil
}
