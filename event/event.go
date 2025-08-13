package event

import "time"

// Event represents an event entity with its properties.
type Event interface {
	GetID() int64                 // Returns the unique identifier of the event
	GetCreatedAt() time.Time      // Returns the creation timestamp of the event
	GetUpdatedAt() time.Time      // Returns the last update timestamp of the event
	GetDeletedAt() time.Time      // Returns the deletion timestamp if the event is currently deleted or 0001-01-01 00:00:00 +0000 UTC if it is not
	GetType() string              // Returns the type of the event i.e create or update or delete
	GetEntity() string            // Returns the entity involved in the event i.e category or product or transaction or user or shop
	GetEntityID() int64           // Returns the primary key of the entity involved in the event i.e the transaction_id
	GetSynchronizedAt() time.Time // Returns the time at which the event was successfully synchronized or 0001-01-01 00:00:00 +0000 UTC if it is not
}
