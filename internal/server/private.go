package server

import (
	http_ext "back-api/pkg/http"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Private struct {
}

func NewPrivateServer() *Private {
	return &Private{}
}

func (s *Private) Run(addr string) {
	r := NewRouter()
	r.PanicHandler = panicHandler

	r.GET("/", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		success := http_ext.NewJsonSuccess(struct{ Status string }{Status: "private server is running"})
		success.Respond(w, req)
	})

	log.Fatalf("server listen and serve error: %s", http.ListenAndServe(addr, r))
}
