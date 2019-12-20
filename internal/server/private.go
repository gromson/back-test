package server

import (
	"back-api/internal/message"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Private struct {
	storage Storage
}

func NewPrivateServer(storage Storage) *Private {
	return &Private{storage: storage}
}

func (s *Private) Run(addr string) {
	r := NewRouter()
	r.PanicHandler = panicHandler

	r.GET("/", privateRootHandle)
	r.GET("/messages", s.fetchHandle())
	r.GET("/messages/:id", s.getMessageHandle())
	r.PUT("/messages/:id", s.updateHandle())


	log.Fatalf("server listen and serve error: %s", http.ListenAndServe(addr, r))
}

func (s *Private) fetchHandle() httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		messages, err := s.storage.Fetch()

		if err != nil {
			log.Println("error while trying to fetch messages from a storage: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		success := NewJsonSuccess(messages)
		success.Respond(w, req)
	}
}

func (s *Private) getMessageHandle() httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		parId := p.ByName("id")
		id, err := uuid.Parse(parId)

		if err != nil {
			problem := NewProblem("Wrong id parameter given", err)
			problem.Respond(w, req)
			return
		}

		msg, err := s.storage.Get(id)

		if err != nil {
			log.Println("error while trying to get a message from a storage: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		success := NewJsonSuccess(msg)
		success.Respond(w, req)
	}
}

func (s *Private) updateHandle() httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		parId := p.ByName("id")
		id, err := uuid.Parse(parId)

		if err != nil {
			problem := NewProblem("Wrong id parameter given", err)
			problem.Respond(w, req)
			return
		}

		msg := new(message.Message)

		if !parsePayloadOrRespond(w, req, msg) {
			return
		}

		if err := s.storage.Edit(id, *msg.Text); err != nil {
			log.Println("error while trying to update a message into a storage: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func privateRootHandle(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	success := NewJsonSuccess(struct{ Status string }{Status: "private server is running"})
	success.Respond(w, req)
}