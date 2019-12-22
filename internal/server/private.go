package server

import (
	"back-api/internal/message"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Authenticator interface {
	Authenticate(email, password string) (uuid.UUID, error)
}

type Private struct {
	storage Storage
	auth Authenticator
}

func NewPrivateServer(storage Storage, auth Authenticator) *Private {
	return &Private{
		storage: storage,
		auth: auth,
	}
}

func (s *Private) Run(addr string) {
	r := NewRouter()
	r.PanicHandler = panicHandler

	auth := r.NewGroup()
	auth.AddMiddleware(s.authMiddleware())

	auth.GET("/", privateRootHandle)
	auth.GET("/messages", s.fetchHandle())
	auth.GET("/messages/:id", s.getMessageHandle())
	auth.PUT("/messages/:id", s.updateHandle())

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

func (s *Private) authMiddleware() Middleware {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
			user, password, ok := req.BasicAuth()

			if !ok {
				basicAuthHeaderFailed(w, req)
				return
			}

			id, err := s.auth.Authenticate(user, password)

			if err != nil {
				authFailed(w, req, err)
				return
			}

			req.Header.Set("User-ID", id.String())

			next(w, req, params)
		}
	}
}

func privateRootHandle(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	success := NewJsonSuccess(struct{ Status string }{Status: "private server is running"})
	success.Respond(w, req)
}

func basicAuthHeaderFailed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"WWW-Authenticate",
		`Basic realm="Wrong data provided", error="invalid_request", error_description="Could not get user and password data"`,
	)
	unauthorized := NewUnauthorizedResponse("Could not get user and password data")
	unauthorized.Respond(w, r)
}

func authFailed(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set(
		"WWW-Authenticate",
		`Basic realm="API", error="invalid_token", error_description="Username and password are not match"`,
	)
	unauthorized := NewUnauthorizedResponse(err.Error())
	unauthorized.Respond(w, r)
}
