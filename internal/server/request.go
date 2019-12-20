package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func parsePayloadOrRespond(w http.ResponseWriter, r *http.Request, s interface{}) {
	payload, err := getPayload(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(payload) == 0 {
		problem := NewProblem(
			"No data provided",
			"The body of the response does not contain the data: "+string(payload),
		)
		problem.Respond(w, r)
		return
	}

	if err := json.Unmarshal(payload, s); err != nil {
		log.Printf("could not unmarshal request body: " + string(payload) + ": " + err.Error())
		detail := ""

		switch e := err.(type) {
		case *json.UnmarshalTypeError:
			detail = fmt.Sprintf("Field %s has a wrong format", e.Field)
		}

		problem := NewProblem("Request payload has an invalid format", detail)
		problem.Respond(w, r)
		return
	}
}

func getPayload(r *http.Request) ([]byte, error) {
	payload, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return payload, fmt.Errorf("could not get the payload: %w", err)
	}

	return payload, nil
}
