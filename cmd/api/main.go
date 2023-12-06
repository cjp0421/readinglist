package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello")
	http.HandleFunc("/v1/healthcheck", healthcheck) //api path? and function to execute when getting incoming requests?
	err := http.ListenAndServe(":4000", nil)        // port and nil - which means it defaults to the go default servemux - problematic because it is not secure
	if err != nil {
		fmt.Println(err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	env := "dev"
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", env)
	fmt.Fprintf(w, "version: %s\n", "1.0.0")
}
