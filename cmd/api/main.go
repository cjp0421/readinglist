package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	flag.Parse()

	fmt.Println("hello")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	addr := fmt.Sprintf(":%d", cfg.port)

	logger.Printf("starting %s server on %s", cfg.env, addr)
	err := http.ListenAndServe(addr, mux)       //using a defined serve mux like this prevents someone else from redefining the global variable used if done as nil, thus making it more secure
	if err != nil {
		fmt.Println(err)
	}
}

// This is the healthcheck endpoint
func healthcheck(w http.ResponseWriter, r *http.Request) {

}
