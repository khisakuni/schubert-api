package api

import (
	"log"
	"net/http"
)

type appHandler func(http.ResponseWriter, *http.Request) *apiError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		log.Printf("Code: %v - %v\n", err.Code, err.Error)
		http.Error(w, err.Message, err.Code)
	}
}

type apiError struct {
	Error   error
	Message string
	Code    int
}

func badInputError(message string) *apiError {
	return &apiError{Code: 400, Message: message, Error: nil}
}

func internalServerError(err error) *apiError {
	return &apiError{Code: 500, Message: "Something went wrong.", Error: err}
}

func notFoundError(err error) *apiError {
	return &apiError{Code: 404, Message: "Record not found", Error: err}
}
