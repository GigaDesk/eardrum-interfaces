package transaction

import (
	"time"
)

// PurchasedProduct represents the essential details of a product bought as part of a transaction.
type PurchasedProduct interface {
	GetProductID() int64 // GetProductID returns the unique identifier of a product purchased in a transaction.
	GetUnitsBought() int // GetUnitsBought returns the number of units of a product purchased.
}

// Transaction represents the core properties of a financial transaction.
type Transaction interface {
	GetID() int64 // GetID returns the unique identifier of the transaction.
	GetUserID() int64 // GetUserID returns the unique identifier of the user who made the transaction.
	GetMerchantID() int64 // GetMerchantID returns the unique identifier of the merchant where the transaction occurred.
	GetCreatedAt() time.Time // GetCreatedAt returns the creation timestamp of the transaction.
	GetUpdatedAt() time.Time // GetUpdatedAt returns the last update timestamp of the transaction.
	GetDeletedAt() time.Time // GetDeletedAt returns the deletion timestamp. The timestamp is a zero value if the transaction has not been deleted.
	GetTotalAmountInCents() uint // GetTotalAmountInCents returns the total amount spent in the transaction, in cents.
	GetTransactionCostInCents() uint // GetTransactionCostInCents returns the transaction processing cost in cents.
}

// TransactionAuthorization provides the data required for a user to authorize a transaction.
// It is intended for embedding in more specific transaction types to avoid code duplication.
type TransactionAuthorization interface {
	GetUUID() string // GetUUID returns the unique UUID string provided by the user to authorize the transaction.
	GetPinCode() string // GetPinCode returns the security PIN code provided by the user for authorization.
}

// NewProductsTransaction represents a new transaction for a set of product purchases.
// It embeds the TransactionAuthorization interface to include common authorization data.
type NewProductsTransaction interface {
	TransactionAuthorization
	GetPurchasedProducts() []PurchasedProduct // GetPurchasedProducts returns information on all product purchases involved in the transaction.
}

// NewTransaction represents a new transaction with a single total amount, not tied to specific products.
// It embeds the TransactionAuthorization interface to include common authorization data.
type NewTransaction interface {
	TransactionAuthorization
	GetTotalAmountInCents() uint // GetTotalAmountInCents returns the total amount intended to be transacted.
}


