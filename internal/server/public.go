package server

import (
	"back-api/internal/server/public"
	http_ext "back-api/pkg/http"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Public struct {
}

func NewPublicServer() *Public {
	return &Public{}
}

func (s *Public) Run(addr string) {
	r := NewRouter()
	r.PanicHandler = panicHandler

	r.GET("/", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		success := http_ext.NewJsonSuccess(struct{ Status string }{Status: "public server is running"})
		success.Respond(w, req)
	})
	r.POST("/message", public.AddMessageHandle())

	log.Fatalf("server listen and serve error: %s", http.ListenAndServe(addr, r))
}
