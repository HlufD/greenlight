package main

import (
	"errors"
	"net/http"
	"strconv"
)

func (app *Application) readParamId(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 0 {

		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}
