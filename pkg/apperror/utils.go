package apperror

import "net/http"

// NewInternalServerError creates a new Internal Server error.
func NewInternalServerError(message string) *AppError {
	return New(message, http.StatusInternalServerError)
}

// NewBadRequest creates a new Bad Request error.
func NewBadRequest(message string) *AppError {
	return New(message, http.StatusBadRequest)
}

// NewUnauthorized creates a new Unauthorized error.
func NewUnauthorized(message string) *AppError {
	return New(message, http.StatusUnauthorized)
}

// IsInternalServerError validates if AppError is from instance of Internal Server Error
func IsInternalServerError(err *AppError) bool {
	return err.http.StatusCode == http.StatusInternalServerError
}

// IsBadRequestError validates if AppError is from instance of Internal Server Error
func IsBadRequestError(err *AppError) bool {
	return err.http.StatusCode == http.StatusBadRequest
}
