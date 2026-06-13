package merchant

import "time"

// Merchant represents a merchant entity with its properties.
type Merchant interface {
	GetID() int64            // Returns the unique identifier of the merchant
	GetCreatedAt() time.Time // Returns the creation timestamp of the merchant
	GetUpdatedAt() time.Time // Returns the last update timestamp of the merchant
	GetDeletedAt() time.Time // Returns the deletion timestamp if the merchant is currently deleted or 0001-01-01 00:00:00 +0000 UTC if it is not
	GetUserName() string     // Returns the username of the merchant
	GetPhoneNumber() string  // Returns the phone number of the merchant
	GetPassword() string     // Returns the security password of the merchant
	GetAccountBalanceInCents() int64 // Returns the merchant's account balance in cents as an integer.
	GetPinCode() *string // Returns the merchant's security pin code
}

// NewMerchant represents data fed into the system with the aim of creating a new merchant entity.
type NewMerchant interface {
	GetUserName() string        // Returns the username of the  new merchant
	GetPhoneNumber() string // Returns the phone number of the new merchant
	GetPassword() string    // Returns the security password of the new merchant
}
