package errors

import (
    "fmt"
    "net/http"
    "github.com/rs/zerolog/log"
)

type SystemCode string

const (
    EARInternalError     SystemCode = "EAR-001"
    EARInvalidInput      SystemCode = "EAR-002"
    EARUserNotFound      SystemCode = "EAR-003" // Legacy / Generic
    EARDuplicateUser     SystemCode = "EAR-004" // Legacy / Generic
    EARUnauthenticated   SystemCode = "EAR-005"
    EARPermissionDenied  SystemCode = "EAR-006"
    EARFileNotFound      SystemCode = "EAR-007"

    // ==========================================
    // MERCHANT SPECIFIC ERRORS (EAR-008 to EAR-019)
    // ==========================================
    EARMerchantInvalidInput             SystemCode = "EAR-008"
    EARMerchantNotFoundByPhone          SystemCode = "EAR-009"
    EARMerchantNotFoundByUsername       SystemCode = "EAR-010"
    EARMerchantPhoneExistsVerified      SystemCode = "EAR-011"
    EARMerchantPhoneExistsUnverified    SystemCode = "EAR-012"
    EARMerchantUsernameExistsVerified   SystemCode = "EAR-013"
    EARMerchantUsernameExistsUnverified  SystemCode = "EAR-014"
    EARMerchantUnauthenticated          SystemCode = "EAR-015"
    EARMerchantPermissionDenied         SystemCode = "EAR-016"
    EARMerchantLookupFailedByPhone      SystemCode = "EAR-017"
    EARMerchantLookupFailedByUsername   SystemCode = "EAR-018"
    EARMerchantListRetrievalFailed      SystemCode = "EAR-019"

    // ==========================================
    // USER SPECIFIC ERRORS (EAR-020 to EAR-031)
    // ==========================================
    EARUserInvalidInput                 SystemCode = "EAR-020"
    
    // Specific User Not Found Scenarios
    EARUserNotFoundByPhone              SystemCode = "EAR-021"
    EARUserNotFoundByUsername           SystemCode = "EAR-022"
    
    // Specific User Conflict Scenarios
    EARUserPhoneExistsVerified          SystemCode = "EAR-023"
    EARUserPhoneExistsUnverified        SystemCode = "EAR-024"
    EARUserUsernameExistsVerified       SystemCode = "EAR-025"
    EARUserUsernameExistsUnverified     SystemCode = "EAR-026"
    
    EARUserUnauthenticated              SystemCode = "EAR-027"
    EARUserPermissionDenied             SystemCode = "EAR-028"

    // User Database Lookup Failures (HTTP 500)
    EARUserLookupFailedByPhone          SystemCode = "EAR-029"
    EARUserLookupFailedByUsername       SystemCode = "EAR-030"
    EARUserListRetrievalFailed          SystemCode = "EAR-031"

    // ==========================================
    // TRANSACTION SPECIFIC ERRORS (EAR-032 onwards)
    // ==========================================
    EARTxMerchantAccountNotFound        SystemCode = "EAR-032"
    EARTxUserAccountNotFound            SystemCode = "EAR-033"
    EARTxInsufficientBalance            SystemCode = "EAR-034"
    EARTxAmountMustBeGreaterThanZero    SystemCode = "EAR-035"
    EARTxInvalidAuthentication          SystemCode = "EAR-036"
)

type errorDefinition struct {
    HttpStatus int
    Message    string
}

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

    // ------------------------------------------
    // Merchant Registrations
    // ------------------------------------------
    EARMerchantInvalidInput: {
        HttpStatus: http.StatusBadRequest,
        Message:    "The provided merchant registration or input details are invalid.",
    },
    EARMerchantNotFoundByPhone: {
        HttpStatus: http.StatusNotFound,
        Message:    "No merchant account matches the provided phone number.",
    },
    EARMerchantNotFoundByUsername: {
        HttpStatus: http.StatusNotFound,
        Message:    "No merchant account matches the provided username.",
    },
    EARMerchantPhoneExistsVerified: {
        HttpStatus: http.StatusConflict,
        Message:    "This phone number is already registered to a verified merchant account.",
    },
    EARMerchantPhoneExistsUnverified: {
        HttpStatus: http.StatusConflict,
        Message:    "This phone number belongs to an unverified merchant. Please complete your verification process.",
    },
    EARMerchantUsernameExistsVerified: {
        HttpStatus: http.StatusConflict,
        Message:    "This username is already taken by a verified merchant.",
    },
    EARMerchantUsernameExistsUnverified: {
        HttpStatus: http.StatusConflict,
        Message:    "This username is reserved by an unverified merchant profile.",
    },
    EARMerchantUnauthenticated: {
        HttpStatus: http.StatusUnauthorized,
        Message:    "Merchant authentication failed or credentials are missing.",
    },
    EARMerchantPermissionDenied: {
        HttpStatus: http.StatusForbidden,
        Message:    "This merchant account does not have permission to access this resource.",
    },
    EARMerchantLookupFailedByPhone: {
        HttpStatus: http.StatusInternalServerError,
        Message:    "The system encountered an error while scanning for the merchant phone number.",
    },
    EARMerchantLookupFailedByUsername: {
        HttpStatus: http.StatusInternalServerError,
        Message:    "The system encountered an error while scanning for the merchant username.",
    },
    EARMerchantListRetrievalFailed: {
        HttpStatus: http.StatusInternalServerError,
        Message:    "The system encountered an error while attempting to compile the merchant list.",
    },

    // ------------------------------------------
    // User Registrations
    // ------------------------------------------
    EARUserInvalidInput: {
        HttpStatus: http.StatusBadRequest,
        Message:    "The provided user input or profile details are invalid.",
    },
    EARUserNotFoundByPhone: {
        HttpStatus: http.StatusNotFound,
        Message:    "No user account matches the provided phone number.",
    },
    EARUserNotFoundByUsername: {
        HttpStatus: http.StatusNotFound,
        Message:    "No user account matches the provided username.",
    },
    EARUserPhoneExistsVerified: {
        HttpStatus: http.StatusConflict,
        Message:    "This phone number is already registered to a verified user account.",
    },
    EARUserPhoneExistsUnverified: {
        HttpStatus: http.StatusConflict,
        Message:    "This phone number belongs to an unverified user. Please complete your verification process.",
    },
    EARUserUsernameExistsVerified: {
        HttpStatus: http.StatusConflict,
        Message:    "This username is already taken by a verified user.",
    },
    EARUserUsernameExistsUnverified: {
        HttpStatus: http.StatusConflict,
        Message:    "This username is reserved by an unverified user profile.",
    },
    EARUserUnauthenticated: {
        HttpStatus: http.StatusUnauthorized,
        Message:    "User authentication failed or credentials are missing.",
    },
    EARUserPermissionDenied: {
        HttpStatus: http.StatusForbidden,
        Message:    "This user account does not have permission to access this resource.",
    },
    EARUserLookupFailedByPhone: {
        HttpStatus: http.StatusInternalServerError,
        Message:    "The system encountered an error while scanning for the user phone number.",
    },
    EARUserLookupFailedByUsername: {
        HttpStatus: http.StatusInternalServerError,
        Message:    "The system encountered an error while scanning for the user username.",
    },
    EARUserListRetrievalFailed: {
        HttpStatus: http.StatusInternalServerError,
        Message:    "The system encountered an error while attempting to compile the user list.",
    },

    // ------------------------------------------
    // Transactions
    // ------------------------------------------
    EARTxMerchantAccountNotFound: {
        HttpStatus: http.StatusNotFound,
        Message:    "The merchant account associated with this transaction could not be found.",
    },
    EARTxUserAccountNotFound: {
        HttpStatus: http.StatusNotFound,
        Message:    "The user account associated with this transaction could not be found.",
    },
    EARTxInsufficientBalance: {
        HttpStatus: http.StatusPaymentRequired,
        Message:    "Transaction declined. The account balance is insufficient to complete this payment.",
    },
    EARTxAmountMustBeGreaterThanZero: {
        HttpStatus: http.StatusBadRequest,
        Message:    "Invalid transaction details. The transaction amount must be greater than zero.",
    },
    EARTxInvalidAuthentication: {
        HttpStatus: http.StatusUnauthorized,
        Message:    "Transaction authentication failed or is invalid.",
    },
}

// PublicError holds the structured details sent back to clients/GraphQL.
type PublicError struct {
    SystemCode SystemCode `json:"code"`        
    HttpStatus int        `json:"http_status"` 
    Message    string     `json:"message"`     
    Err        error      `json:"-"`           
}

func (e *PublicError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("[%s] (HTTP %d) %s (wrapped: %s)", e.SystemCode, e.HttpStatus, e.Message, e.Err.Error())
    }
    return fmt.Sprintf("[%s] (HTTP %d) %s", e.SystemCode, e.HttpStatus, e.Message)
}

func (e *PublicError) Unwrap() error {
    return e.Err
}

func New(code SystemCode, err error) *PublicError {
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

func (e *PublicError) Log() {
    if e.HttpStatus >= 500 {
        log.Error().
            Str("code", string(e.SystemCode)).
            Int("status", e.HttpStatus).
            Err(e.Err).
            Msg(e.Message)
    } else {
        log.Warn().
            Str("code", string(e.SystemCode)).
            Int("status", e.HttpStatus).
            Msg(e.Message)
    }
}