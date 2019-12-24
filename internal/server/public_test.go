package server

import (
	"back-api/internal/message"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type storageTest map[uuid.UUID]*message.Message

// Get retrieves a message from a storage by its id
func (s storageTest) Get(id uuid.UUID) (*message.Message, error) {
	if m, ok := s[id]; ok {
		return m, nil
	}

	return nil, errors.New("message not found")
}

// Fetch fetches all messages from a storage. Latest messages from in the beginning
func (s storageTest) Fetch() ([]*message.Message, error) {
	res := make([]*message.Message, 0, len(s))

	for _, m := range s {
		res = append(res, m)
	}

	return res, nil
}

// Insert inserts new message into a storage
func (s storageTest) Insert(m *message.Message) error {
	s[*m.ID] = m
	return nil
}

// Edit modifies the message by given ID. Nil fields will be ignored
func (s storageTest) Edit(id uuid.UUID, text string) error {
	s[id].Text = &text
	return nil
}

type successfulResponseTest struct {
	ID uuid.UUID `json:"id"`
}

var testCases = []struct {
	data string
	code int
}{
	{
		`{"name": "Roman Iudin", "email": "gromson@gmail.com", "text": "Successfully added message"}`,
		http.StatusOK,
	},
	{
		`{"name": "Perdo Gonsalez", "email": "pg@mail.com"}`,
		http.StatusBadRequest,
	},
}

func TestPublic_addMessageHandle(t *testing.T) {
	server := NewPublicServer(storageTest{})
	handle := server.addMessageHandle()

	for _, c := range testCases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/messages", strings.NewReader(c.data))
		handle(w, req, nil)

		if w.Code != c.code {
			t.Errorf("request %s expected to get %d code, %d received", c.data, c.code, w.Code)
		}

		if w.Code == http.StatusOK {
			res := new(successfulResponseTest)
			err := json.Unmarshal(w.Body.Bytes(), res)

			if err != nil {
				t.Error("couldn't read a response body")
			}

			if reflect.DeepEqual(res.ID, uuid.UUID{}) {
				t.Error("empty ID returned")
			}
		}
	}

}
