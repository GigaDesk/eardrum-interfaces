package purchase

// Purchase represents a purchase entity with its properties.
type Purchase interface {
	GetID() int64            // Returns the unique identifier of the purchase
	GetTransactionID() int64 // Returns the unique identifier of the transaction in which the purchase was made
	GetProductID() int64     // Returns the unique identifier of a product purchased
	GetUnitsBought() int     // Returns the the amount of units of a product purchased
	GetTotalAmountInCents() int64   // Returns the total amount spent in the purchase in cents
}
