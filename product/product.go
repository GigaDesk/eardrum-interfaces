package product

import "time"

// Product represents a product entity with its properties.
type Product interface {
	GetID() int64                  // Returns the unique identifier of the product
	GetCreatedAt() time.Time       // Returns the creation timestamp of the product
	GetUpdatedAt() time.Time       // Returns the last update timestamp of the product
	GetDeletedAt() time.Time       // Returns the deletion timestamp if the product is currently deleted or 0001-01-01 00:00:00 +0000 UTC if it is not
	GetName() string               // Returns the name of the product
	GetPricePerUnitInCents() int64 // Returns the product's price per unit in cents as an integer.
}

// NewProduct represents data fed into the system with the aim of creating a new product entity.
type NewProduct interface {
	GetName() string               // Returns the name of the product
	GetPricePerUnitInCents() int64 // Returns the product's price per unit in cents as an integer.
}
