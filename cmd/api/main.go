package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// this is config for the app ,like port and environment
type config struct {
	port int
	env  string
}

// This is a struct that holds dependencies for the application like  database, for now only config and logger
type Application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "This the port on which the app runs")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		config: cfg,
		logger: logger,
	}

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.port),
		Handler:           app.routes(),
		ReadTimeout:       10 * time.Second,
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 30 * time.Second,
	}

	logger.Printf("The server is running on port: %d", cfg.port)

	err := server.ListenAndServe()

	if err != nil {
		logger.Fatal(err)
	}
}
