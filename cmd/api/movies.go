package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/greenlight/internal/data"
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

	movie := data.Movie{
		ID:    id,
		Title: "",
		//Year:      2022,
		Runtime:   104,
		Genres:    []string{"drama", "romance", "war"},
		CreatedAt: time.Now(),
		Version:   version,
	}

	//fmt.Fprintf(w, "Got the movie with %d", id)
	err = app.writeJson(w, 200, envelop{"movie": movie}, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
