package errors 

import (
	"fmt"
	"net/http"
)

// PublicError holds structured details (Code, Reason) for GraphQL clients.
type PublicError struct {
	Message string // The public message for the client
	Code    int    // The HTTP status code hint (e.g., 400, 500)
	Reason  string // A machine-readable string (e.g., "BAD_INPUT", "CONFLICT")
	Err     error  // The underlying private error (for logging/debugging)
}

// Implement the Go error interface
func (e *PublicError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %s (wrapped: %s)", e.Code, e.Reason, e.Message, e.Err.Error())
	}
	return fmt.Sprintf("[%d] %s: %s", e.Code, e.Reason, e.Message)
}

// Implement the Unwrap method for errors.Is/As to work correctly
func (e *PublicError) Unwrap() error {
	return e.Err
}

// NewHTTPError is the primary way to create a structured error
func NewHTTPError(status int, message string, err error) *PublicError {
	return &PublicError{
		Message: message,
		Code:    status,
		Reason:  http.StatusText(status), // Using standard HTTP text is simple
		Err:     err,
	}
}