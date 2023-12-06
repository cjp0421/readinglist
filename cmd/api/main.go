package main

import (
	"flag"
	"fmt"
	"net/http"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

func main() {
	var cfg config

	flag.IntVar()

	fmt.Println("hello")

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", healthcheck) //api endpoint and function to execute when getting incoming requests?
	err := http.ListenAndServe(":4000", mux)       //using a defined serve mux like this prevents someone else from redefining the global variable used if done as nil, thus making it more secure
	if err != nil {
		fmt.Println(err)
	}
}

// This is the healthcheck endpoint
func healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// env := "dev"

	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", env)
	fmt.Fprintf(w, "version: %s\n", version)
}
