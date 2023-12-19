package main

import (
	"encoding/json"
	"net/http"
)

// the type below is part of making an envelope for JSON data
// this envelope type will be used to collect the JSON data within a named object which can make parsing easier
type envelope map[string]any

// Credit: Alex Edwards, Let's Go Further
// This was added to replace duplicated code in the handlers so as to observed the DRY principle
func (app *application) writeJSON(w http.ResponseWriter, status int, data any) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return http.ErrAbortHandler
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
