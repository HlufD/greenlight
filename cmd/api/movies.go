package main

import (
	"fmt"
	"net/http"
)

func (app *Application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "movie created")
}

func (app *Application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// get the id from  the path
	id, err := app.readParamId(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Got the movie with %d", id)
}
