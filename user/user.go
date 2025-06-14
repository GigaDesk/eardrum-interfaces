package student

import "time"

// User represents a user entity with its properties.
type User interface {
	GetID() int64              // Returns the unique identifier of the user
	GetCreatedAt() time.Time // Returns the creation timestamp of the user
	GetUpdatedAt() time.Time // Returns the last update timestamp of the user
	GetDeletedAt() time.Time // Returns the deletion timestamp if the user is currently deleted or 0001-01-01 00:00:00 +0000 UTC if not
	GetName() string         // Returns the name of the user
	GetPhoneNumber() string  // Returns the phone number of the user
	GetPassword() string     // Returns the security password of the user
	GetAccountBalanceInCents() int64 // Returns the user's account balance in cents as an integer.
	GetPinCode() string // Returns the user's security pin code
}

// NewUser represents data fed into the system with the aim of creating a new user entity.
type NewUser interface {
	GetName() string         // Returns the name of the user
	GetPhoneNumber() string  // Returns the phone number of the user
	GetPassword() string     // Returns the security password of the user
}