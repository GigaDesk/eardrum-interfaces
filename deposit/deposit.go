package deposit

import "time"

// Deposit represents a deposit entity with its properties.
type Deposit interface {
	GetID() int64               // Returns the unique identifier of the user
	GetUserID() int64           // Returns the unique identifier of the user that the deposit was directed to
	GetCreatedAt() time.Time    // Returns the creation timestamp of the user
	GetUpdatedAt() time.Time    // Returns the last update timestamp of the user
	GetDeletedAt() time.Time    // Returns the deletion timestamp if the user is currently deleted or 0001-01-01 00:00:00 +0000 UTC if not
	GetCode() string            // Returns the unique mpesa transaction code
	GetAmount() int64           // Returns the deposit amount i.e amount deposited.
	GetDepositerNumber() string // Returns the number of the depositer
}

// NewDeposit represents data fed into the system with the aim of creating a new deposit entity.
type NewDeposit interface {
	GetAmount() int64           // Returns the deposit amount i.e amount deposited.
	GetUserPhoneNumber() string // Returns the phone number of the user
	GetCode() string            // Returns the unique mpesa transaction code
	GetDepositerNumber() string // Returns the number of the depositer
}
