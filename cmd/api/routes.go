package main

import "net/http"

func (app *Application) routes() *http.ServeMux {
	// initialize a touter
	router := http.NewServeMux()

	// define a health check route with method and path
	router.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)
	router.HandleFunc("POST /v1/movies", app.createMovieHandler)
	router.HandleFunc("GET /v1/movies/{id}", app.showMovieHandler)
	return router
}
