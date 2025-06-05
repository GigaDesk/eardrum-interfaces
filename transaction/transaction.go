package transaction

import "time"

// PurchasedProduct represents basic data fed into the system about the purchase of a particular product
type PurchasedProduct interface {
	GetProductID() int64 // Returns the unique identifier of a product purchased in a transaction
	GetUnitsBought() int // Returns the the amount of units of a product purchased in a transaction
}

// Transaction represents a transaction entity with its properties.
type Transaction interface {
	GetID() int64            // Returns the unique identifier of the transaction
	GetCreatedAt() time.Time // Returns the creation timestamp of the transaction
	GetUpdatedAt() time.Time // Returns the last update timestamp of the transaction
	GetDeletedAt() time.Time // Returns the deletion timestamp if the transaction is currently deleted or 0001-01-01 00:00:00 +0000 UTC if it is not
	GetTotalAmountInCents() int64   // Returns the total amount spent in the transaction in cents
	GetBalanceBeforeInCents() int64 // Returns the account balance before the transaction in cents
	GetBalanceAfterInCents() int64  // Returns the account balance after the transaction in cents.
}

// NewTransaction represents data fed into the system with the aim of creating a new transaction entity.
type NewTransaction interface {
	GetPurchasedProducts() []PurchasedProduct //Returns information on all product purchases involved in the transaction
	GetRegistrationNumber() string//Returns the registration number provided by the student to authorize the transaction
	GetPinCode() string//Returns the pin code provided by the student to authorize the transaction
}
