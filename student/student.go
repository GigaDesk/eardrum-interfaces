package student

import "time"

// Student represents a student entity with its properties.
type Student interface {
	GetID() int64              // Returns the unique identifier of the student
	GetCreatedAt() time.Time // Returns the creation timestamp of the student
	GetUpdatedAt() time.Time // Returns the last update timestamp of the student
	GetDeletedAt() time.Time // Returns the deletion timestamp if the student is currently deleted or 0001-01-01 00:00:00 +0000 UTC if not
	GetRegistrationNumber() string // Returns the student's registration number
	GetName() string         // Returns the name of the student
	GetPhoneNumber() string  // Returns the phone number of the student
	GetDateOfAdmission()  time.Time       // Returns the student's date of admission
	GetDateofBirth() time.Time     // Returns the student's birthday
	GetProfilePicture() string  //Returns the student profile picture's image url
	GetPassword() string     // Returns the security password of the student
	GetAccountBalanceInCents() int64 // Returns the student's account balance in cents as an integer.
	GetPinCode() string // Returns the student's security pin code
}

// NewStudent represents data fed into the system with the aim of creating a new student entity.
type NewStudent interface {
	GetRegistrationNumber() string // Returns the student's registration number
	GetName() string         // Returns the name of the student
	GetPhoneNumber() string  // Returns the phone number of the student
	GetDateOfAdmission()  time.Time       // Returns the student's date of admission
	GetDateofBirth() time.Time     // Returns the student's birthday
	GetProfilePicture() string  //Returns the student profile picture's image url
	GetPassword() string     // Returns the security password of the student
}