package server

import (
	"back-api/internal/message"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type Public struct {
	storage Storage
}

func NewPublicServer(storage Storage) *Public {
	return &Public{storage: storage}
}

func (s *Public) Run(addr string) {
	r := NewRouter()
	r.PanicHandler = panicHandler

	r.GET("/", publicRootHandle)
	r.POST("/messages", s.addMessageHandle())

	log.Fatalf("server listen and serve error: %s", http.ListenAndServe(addr, r))
}

func (s *Public) addMessageHandle() httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		msg := &message.Message{}

		if !parsePayloadOrRespond(w, req, msg) {
			return
		}

		generatedUUID := uuid.New()
		msg.ID = &generatedUUID

		creationTime := time.Now()
		msg.CreationTime = &creationTime

		errs := validateMessage(msg)

		if len(errs) > 0 {
			problem := NewProblem("Wrong data provided", errs)
			problem.Respond(w, req)
			return
		}

		if err := s.storage.Insert(msg); err != nil {
			log.Println("error while trying to insert a message into a storage: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		success := NewJsonSuccess(
			struct {
				ID *uuid.UUID `json:"uuid"`
			}{
				ID: msg.ID,
			},
		)
		success.Respond(w, req)
	}
}

func publicRootHandle(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	success := NewJsonSuccess(struct{ Status string }{Status: "public server is running"})
	success.Respond(w, req)
}
