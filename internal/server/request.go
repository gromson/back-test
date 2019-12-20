package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// parsePayloadOrRespond tries to parse a payload and populate given s interface
func parsePayloadOrRespond(w http.ResponseWriter, r *http.Request, s interface{}) bool {
	payload, err := getPayload(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}

	if len(payload) == 0 {
		problem := NewProblem(
			"No data provided",
			"The body of the response does not contain the data: "+string(payload),
		)
		problem.Respond(w, r)
		return false
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
		return false
	}

	return true
}

// getPayload returns request payload as a byte slice
func getPayload(r *http.Request) ([]byte, error) {
	payload, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return payload, fmt.Errorf("could not get the payload: %w", err)
	}

	return payload, nil
}
