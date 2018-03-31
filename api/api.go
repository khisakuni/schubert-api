package api

import "net/http"

type appHandler func(http.ResponseWriter, *http.Request) *apiError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Message, err.Code)
	}
}

type apiError struct {
	Error   error
	Message string
	Code    int
}

func internalServerError(err error) *apiError {
	return &apiError{Code: 500, Message: "Something went wrong.", Error: err}
}
