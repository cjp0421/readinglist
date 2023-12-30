package main

import (
	"fmt"
	"net/http"
)

// defines the receiver
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The home page")
}

func (app *application) bookView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "View a single book")
}

func (app *application) bookCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creates a new book record form")
}
