package school

import "time"

// School represents a school entity with its properties.
type School interface {
	GetID() int64            // Returns the unique identifier of the school
	GetCreatedAt() time.Time // Returns the creation timestamp of the school
	GetUpdatedAt() time.Time // Returns the last update timestamp of the school
	GetDeletedAt() time.Time // Returns the deletion timestamp if the school is currently deleted or 0001-01-01 00:00:00 +0000 UTC if it is not
	GetName() string         // Returns the name of the school
	GetPhoneNumber() string  // Returns the phone number of the school
	GetBadge() string        // Returns a badge or identifier associated with the school
	GetWebsite() string      // Returns the website URL of the school
	GetPassword() string     // Returns the security password of the school
}

// NewSchool represents data fed into the system with the aim of creating a new school entity.
type NewSchool interface {
	GetName() string        // Returns the name of the  new school
	GetPhoneNumber() string // Returns the phone number of the new school
	GetBadge() string       // Returns a badge or identifier associated with the new school
	GetWebsite() string     // Returns the website URL of the new school
	GetPassword() string    // Returns the security password of the new school
}
