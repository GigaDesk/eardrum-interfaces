package errors

import (
	"fmt"
	"net/http"
	"github.com/rs/zerolog/log"
)

// SystemCode represents our custom Eardrum error code format (e.g., EAR-001)
type SystemCode string

// Define your Oracle-style system error codes here
const (
	EARInternalError    SystemCode = "EAR-001"
	EARInvalidInput     SystemCode = "EAR-002"
	EARUserNotFound     SystemCode = "EAR-003"
	EARDuplicateUser    SystemCode = "EAR-004"
	EARUnauthenticated  SystemCode = "EAR-005"
	EARPermissionDenied SystemCode = "EAR-006"
	EARFileNotFound     SystemCode = "EAR-007"
)

// errorDefinition binds our internal system code to its default metadata
type errorDefinition struct {
	HttpStatus int
	Message    string
}

// The global registry mapping your EAR codes to default HTTP statuses and messages.
// This serves as your single source of truth for error metadata.
var registry = map[SystemCode]errorDefinition{
	EARInternalError: {
		HttpStatus: http.StatusInternalServerError,
		Message:    "An unexpected internal server error occurred. Please try again later.",
	},
	EARInvalidInput: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The provided input parameters are invalid or malformed.",
	},
	EARUserNotFound: {
		HttpStatus: http.StatusNotFound,
		Message:    "The requested user record could not be found in the system.",
	},
	EARDuplicateUser: {
		HttpStatus: http.StatusConflict,
		Message:    "A user account with those details already exists.",
	},
	EARUnauthenticated: {
		HttpStatus: http.StatusUnauthorized,
		Message:    "Authentication is required to access this resource.",
	},
	EARPermissionDenied: {
		HttpStatus: http.StatusForbidden,
		Message:    "You do not have the required permissions to perform this action.",
	},
	EARFileNotFound: {
		HttpStatus: http.StatusNotFound,
		Message:    "The requested file or asset could not be found.",
	},
}

// PublicError holds the structured details sent back to clients/GraphQL.
type PublicError struct {
	SystemCode SystemCode `json:"code"`        // e.g., "EAR-004"
	HttpStatus int        `json:"http_status"` // e.g., 409
	Message    string     `json:"message"`     // Centralized client-safe message
	Err        error      `json:"-"`           // Hidden internal debugging error (excluded from JSON)
}

// Implement the Go error interface
func (e *PublicError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] (HTTP %d) %s (wrapped: %s)", e.SystemCode, e.HttpStatus, e.Message, e.Err.Error())
	}
	return fmt.Sprintf("[%s] (HTTP %d) %s", e.SystemCode, e.HttpStatus, e.Message)
}

// Implement the Unwrap method for errors.Is/As to work correctly
func (e *PublicError) Unwrap() error {
	return e.Err
}

// New creates a PublicError by automatically looking up the code in our registry.
// You no longer need to explicitly pass messages or HTTP statuses here.
func New(code SystemCode, err error) *PublicError {
	// Fallback in case an undefined code is passed during development
	defn, exists := registry[code]
	if !exists {
		defn = registry[EARInternalError]
		code = EARInternalError
	}

	return &PublicError{
		SystemCode: code,
		HttpStatus: defn.HttpStatus,
		Message:    defn.Message,
		Err:        err,
	}
}



// Log outputs the error cleanly to your server console using zerolog
func (e *PublicError) Log() {
	if e.HttpStatus >= 500 {
		// Severe server crash: log as an Error and include the raw engineering error
		log.Error().
			Str("code", string(e.SystemCode)).
			Int("status", e.HttpStatus).
			Err(e.Err).
			Msg(e.Message)
	} else {
		// Client/User mistake (400, 401, etc.): log as a Warn and skip raw system logs
		log.Warn().
			Str("code", string(e.SystemCode)).
			Int("status", e.HttpStatus).
			Msg(e.Message)
	}
}