package main

import (
	"database/sql" //package provides a generic api that allows for interacting with the databases in a vendor-neutral way
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq" //This is a driver; this is the go package for the sql database driver; third-party package

	"readinglist/internal/data"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	dsn  string // short for data name service; aka a data connection string; this will be passed in so we can connect to the database
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	flag.StringVar(&cfg.dsn, "db-dsn", os.Getenv("READINGLIST_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	fmt.Println("hello")

	cfg.dsn = "postgres://postgres:mysecretpassword@localhost/readinglist?sslmode=disable"
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//below opens the database connection
	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		logger.Fatal(err)
	}

	err = db.Ping() //this tests the connection
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close() //this closes the connection

	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	addr := fmt.Sprintf(":%d", cfg.port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
