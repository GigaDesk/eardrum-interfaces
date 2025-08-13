package withdrawal

import "time"

// Withdrawal represents a withdrawal entity with its properties.
type Withdrawal interface {
	GetID() int64               // Returns the unique identifier of the withdrawal record
	GetUserID() int64           // Returns the unique identifier of the user that the withdrawal was directed to
	GetCreatedAt() time.Time    // Returns the creation timestamp of the withdrawal
	GetUpdatedAt() time.Time    // Returns the last update timestamp of the withdrawal
	GetDeletedAt() time.Time    // Returns the deletion timestamp if the withdrawal is currently deleted or 0001-01-01 00:00:00 +0000 UTC if not
	GetCode() string            // Returns the unique mpesa transaction code from the withdrawal
	GetAmount() int64           // Returns the withdrawal amount i.e amount withdrawed.
	GetMpesaNumber() string     // Returns the m-pesa number that the withdrawal should be disbursed to 
	GetProcessedAt() time.Time  // Returns the timestamp at which the withdrawal was processed
}