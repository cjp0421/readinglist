package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	//these are the routes
	mux.HandleFunc("/", app.home) //if one comes in on the slash http address, the route goes to the app.home page
	mux.HandleFunc("/book/view", app.bookView)
	mux.HandleFunc("/book/create", app.bookCreate)

	return mux
}
