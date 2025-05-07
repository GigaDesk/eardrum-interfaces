package shop

import "time"

// Shop represents a shop entity with its properties.
type Shop interface {
	GetID() int64            // Returns the unique identifier of the shop
	GetCreatedAt() time.Time // Returns the creation timestamp of the shop
	GetUpdatedAt() time.Time // Returns the last update timestamp of the shop
	GetDeletedAt() time.Time // Returns the deletion timestamp if the shop is currently deleted or 0001-01-01 00:00:00 +0000 UTC if it is not
	GetName() string         // Returns the name of the shop
	GetPhoneNumber() string  // Returns the phone number of the shop
	GetPassword() string     // Returns the security password of the shop
	GetAccountBalanceInCents() int64 // Returns the shop's account balance in cents as an integer.
	GetPinCode() string // Returns the shop's security pin code
}

// NewShop represents data fed into the system with the aim of creating a new shop entity.
type NewSchool interface {
	GetName() string        // Returns the name of the  new shop
	GetPhoneNumber() string // Returns the phone number of the new shop
	GetPassword() string    // Returns the security password of the new shop
}
