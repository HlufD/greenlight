package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func (app *Application) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {

		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}

	}

	return nil

}
