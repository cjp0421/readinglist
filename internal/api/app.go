package api

import (
	"log"
	"net/http"
	"readinglist/internal/data"
)

const version = "2.0.0"

type Config struct {
	Port int
	Env  string
	Dsn  string // short for data name service; aka a data connection string; this will be passed in so we can connect to the database
}

type Application struct {
	Config Config
	Logger *log.Logger
	Models data.Models
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
