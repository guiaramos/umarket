package apperror

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fastbill/go-httperrors"
)

// AppError represents the app error handler.
type AppError struct {
	err  error
	http httperrors.HTTPError
}

// New creates a new AppError.
func New(message string, status int) *AppError {
	return &AppError{
		err:  errors.New(message),
		http: *httperrors.New(status, message),
	}
}

// Error returns the string representation of the error message.
func (e AppError) Error() string {
	return e.err.Error()
}

// Unwrap returns the error.
func (e AppError) Unwrap() error {
	fmt.Println(e.err.Error())
	return e.err
}

// HttpError allows to write the http error to a ResponseWriter for implemantation of the WritableError interface
func (e AppError) HttpError(w http.ResponseWriter) error {
	return e.http.WriteJSON(w)
}
