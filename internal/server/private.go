package server

import (
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

	r.GET("/", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		success := NewJsonSuccess(struct{ Status string }{Status: "private server is running"})
		success.Respond(w, req)
	})
	r.GET("/message/:id", s.getMessageHandle())

	log.Fatalf("server listen and serve error: %s", http.ListenAndServe(addr, r))
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
