package main

import (
	"database/sql" //package provides a generic api that allows for interacting with the databases in a vendor-neutral way
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //This is a driver; this is the go package for the sql database driver; third-party package

	"readinglist/internal/api"
	"readinglist/internal/data"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg api.Config
	// dfdsfdsf

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	cfg.Dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	flag.IntVar(&cfg.Port, "port", 3001, "API server port")
	flag.StringVar(&cfg.Env, "env", "dev", "Environment (dev|stage|prod)")
	flag.Parse()

	fmt.Println("hello")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//below opens the database connection
	db, err := sql.Open("postgres", cfg.Dsn)
	if err != nil {
		logger.Fatal(err)
	}

	err = db.Ping() //this tests the connection
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close() //this closes the connection

	logger.Printf("database connection pool established")

	app := &api.Application{
		Config: cfg,
		Logger: logger,
		Models: data.NewModels(db),
	}

	addr := fmt.Sprintf(":%d", cfg.Port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.Route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.Env, addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
