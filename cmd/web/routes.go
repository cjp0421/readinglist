package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	//below adds the file server to the web application, which is needed to access the css
	//when this handler receives a request it removes the leading slash from the url path
	//then it searches the /ui/static directory for the corresponding file to send the director
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	//for the above to work correctly, we have to strip off the leading slash to ensure one is sent to the correct place
	//the StripPrefix removes the leading slash
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//these are the routes
	mux.HandleFunc("/", app.home) //if one comes in on the slash http address, the route goes to the app.home page
	mux.HandleFunc("/book/view", app.bookView)
	mux.HandleFunc("/book/create", app.bookCreate)

	return mux
}
