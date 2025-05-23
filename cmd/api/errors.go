package main

import (
	"fmt"
	"net/http"
)

func (app *Application) errorLog(r *http.Request, err error) {
	app.logger.Println(err)
}

func (app *Application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {

	envelopedResponse := envelop{
		"error": message,
	}

	err := app.writeJson(w, status, envelopedResponse, nil)

	if err != nil {
		app.errorLog(r, err)
		w.WriteHeader(500)
	}
}

func (app *Application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorLog(r, err)

	message := "the server encountered a problem and could not process your request"

	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *Application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"

	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *Application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
