package main

import (
	"net/http"
)

func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	// this is only good for static response
	/*
		js := `{"status":"Available","environment": %q,"version": %q}`
		js = fmt.Sprintf(js, app.config.env, version)
		w.Header().Set("Content-type", "application/json")
		w.Write([]byte(js))
	*/

	// using json.marshal()
	data := map[string]string{
		"status":      "Available",
		"environment": app.config.env,
		"version":     version,
	}
	err := app.writeJson(w, 200, data, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	// w.Header().Set("Content-type", "application/json")
	// w.Write(js)
}
