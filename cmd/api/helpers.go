package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type envelop map[string]interface{}

func (app *Application) readParamId(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 0 {

		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *Application) writeJson(w http.ResponseWriter, status int, data envelop, headers http.Header) error {

	js, err := json.MarshalIndent(data, "", "\t")

	js = append(js, '\n')

	if err != nil {
		return err
	}

	for key, value := range headers {
		// while is this good for custom headers and when setting header with multiple values
		w.Header()[key] = value
	}

	// .Set() is good for standard headers and single headers
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
