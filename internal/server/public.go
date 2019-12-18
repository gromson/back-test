package server

import (
	"back-api/internal/server/public"
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

	r.POST("/message", public.AddMessageHandle())

	log.Println("Public server started on " + addr)
	log.Fatalf("server listen and serve error: %s", http.ListenAndServe(addr, r))
}