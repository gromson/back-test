package authentication

import (
	"back-api/internal/password"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

var testProfiles = testAuthStorage{
	"user1@mail.com": {
		id:       mustUuidParse("aa5cc407-63db-414a-9544-471725a0ac87"),
		password: mustGenerateTestPassword("123"),
	},
	"user2@mail.com": {
		id:       mustUuidParse("be4d8ca7-1b68-44c4-9230-d2528060e976"),
		password: mustGenerateTestPassword("1234"),
	},
	"user3@mail.com": {
		id:       mustUuidParse("b4c90ccb-8549-4e9e-b1fe-376b6681e40d"),
		password: mustGenerateTestPassword("12345"),
	},
}

type testCase struct {
	input  [2]string
	result uuid.UUID
	err    error
}

var testCases = []testCase{
	{
		input:  [2]string{"user1@mail.com", "123"},
		result: mustUuidParse("aa5cc407-63db-414a-9544-471725a0ac87"),
		err:    nil,
	},
	{
		input:  [2]string{"user2@mail.com", "1234"},
		result: mustUuidParse("be4d8ca7-1b68-44c4-9230-d2528060e976"),
		err:    nil,
	},
	{
		input:  [2]string{"user3@mail.com", "1234"},
		result: mustUuidParse("be4d8ca7-1b68-44c4-9230-d2528060e976"),
		err:    ErrInvalidPassword{},
	},
	{
		input:  [2]string{"non-existing@mail.com", "1234"},
		result: uuid.UUID{},
		err:    ErrClientNotFound{},
	},
}

func TestAuth_Authenticate(t *testing.T) {
	auth := NewAuth(&testProfiles, &password.BcryptService{})

	for _, c := range testCases {
		id, err := auth.Authenticate(c.input[0], c.input[1])

		var invalidPasswordError ErrInvalidPassword
		if errors.As(c.err, &invalidPasswordError) && !errors.As(err, &invalidPasswordError) {
			t.Errorf("ErrInvalidPassword error expected, %T given", err)
		}

		if errors.As(c.err, &invalidPasswordError) && errors.As(err, &invalidPasswordError) {
			continue
		}

		var clientNotFoundError ErrClientNotFound
		if errors.As(c.err, &clientNotFoundError) && !errors.As(err, &clientNotFoundError) {
			t.Errorf("ErrInvalidPassword error expected, %T given", err)
		}

		if errors.As(c.err, &clientNotFoundError) && errors.As(err, &clientNotFoundError) {
			continue
		}

		if err != nil {
			t.Error(err)
		}

		if id != c.result {
			t.Errorf("unexpected result: %s expected, %s given", c.result, id)
		}
	}
}

func mustGenerateTestPassword(pswd string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(pswd), 12)

	if err != nil {
		log.Fatal(err)
	}

	return hash
}

func mustUuidParse(s string) uuid.UUID {
	id, err := uuid.Parse(s)

	if err != nil {
		log.Fatal(err)
	}

	return id
}
