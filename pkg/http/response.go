package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response interface {
	Respond(w http.ResponseWriter, r *http.Request)
}

type jsonSuccess struct {
	result interface{}
}

func NewJsonSuccess(result interface{}) Response {
	return &jsonSuccess{result: result}
}

func (r *jsonSuccess) Respond(w http.ResponseWriter, req *http.Request) {
	serialized, err := r.serialize()

	if err != nil {
		log.Println("error while trying to serialize a response", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(serialized); err != nil {
		log.Println("error while trying to write json response body", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (r *jsonSuccess) serialize() ([]byte, error) {
	return json.Marshal(r.result)
}