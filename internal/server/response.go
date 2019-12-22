package server

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
	serialized, err := json.Marshal(r.result)

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

type problem struct {
	Type   string      `json:"type"`
	Title  string      `json:"title"`
	Status int         `json:"status"`
	Detail interface{} `json:"details"`
}

func NewProblem(title string, detail interface{}) Response {
	if errs, ok := detail.([]error); ok {
		detail = errorSliceToStringSlice(errs)
	}

	return &problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.1",
		Title:  title,
		Status: http.StatusBadRequest,
		Detail: detail,
	}
}

func (r *problem) Respond(w http.ResponseWriter, req *http.Request) {
	serialized, err := json.Marshal(r)

	if err != nil {
		log.Printf("error while trying to serialize a problem response: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")

	w.WriteHeader(r.Status)

	if _, err := w.Write(serialized); err != nil {
		log.Printf("error while trying to write problem response body: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

type unauthorized struct {
	problem
}

func NewUnauthorizedResponse(detail string) Response {
	problem := problem{
		Type:   "https://tools.ietf.org/html/rfc7235#section-3.1",
		Title:  "Unauthorized",
		Status: http.StatusUnauthorized,
		Detail: detail,
	}

	return &unauthorized{
		problem,
	}
}

func (r *unauthorized) Respond(w http.ResponseWriter, req *http.Request) {
	r.problem.Respond(w, req)
}

func errorSliceToStringSlice(data []error) []string {
	ss := make([]string, 0, len(data))

	for _, e := range data {
		ss = append(ss, e.Error())
	}

	return ss
}
