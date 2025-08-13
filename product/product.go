package product

import "time"

// Product represents a product entity with its properties.
type Product interface {
	GetID() int64                  // Returns the unique identifier of the product
	GetShopID() int64              // Returns the unique identifier of the shop that the product belongs to
	GetCategoryID() int64          // Returns the unique identifier of the category that the product belongs to
	GetCreatedAt() time.Time       // Returns the creation timestamp of the product
	GetUpdatedAt() time.Time       // Returns the last update timestamp of the product
	GetDeletedAt() time.Time       // Returns the deletion timestamp if the product is currently deleted or 0001-01-01 00:00:00 +0000 UTC if it is not
	GetName() string               // Returns the name of the product
	GetPricePerUnitInCents() int64 // Returns the product's price per unit in cents as an integer.
	GetDeleted() bool              //Returns if the product has been deleted or not i.e true if deleted
	GetBlocked() bool              //Returns if the product has been blocked or not i.e true if blocked
}

// NewProduct represents data fed into the system with the aim of creating a new product entity.
type NewProduct interface {
	GetName() string               // Returns the name of the product
	GetPricePerUnitInCents() int64 // Returns the product's price per unit in cents as an integer.
}

// Category represents a group of product entities with its properties.
type Category interface {
	GetID() int64            // Returns the unique identifier of the category
	GetShopID() int64        // Returns the unique identifier of the shop that the category belongs to
	GetCreatedAt() time.Time // Returns the creation timestamp of the category
	GetUpdatedAt() time.Time // Returns the last update timestamp of the category
	GetDeletedAt() time.Time // Returns the deletion timestamp if the category is currently deleted or 0001-01-01 00:00:00 +0000 UTC if it is not
	GetName() string         // Returns the name of the category
	GetDescription() string  // Returns the category's description
	GetBlocked() bool        //Returns if the category has been blocked or not i.e true if blocked
}

// NewCategory represents data fed into the system with the aim of creating a new Category entity.
type NewCategory interface {
	GetName() string        // Returns the name of the category
	GetDescription() string // Returns the category's description
}
