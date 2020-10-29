package exchanges

import (
	"time"

	"github.com/xmn-services/rod-network/libs/hash"
)

// Application represents an exchange application
type Application interface {
	New(method hash.Hash, from hash.Hash, to hash.Hash, expireOn time.Time, notes string) error
	Update(hash hash.Hash, update Update) error
	Delete(hash hash.Hash) error
}

// Update represents an update
type Update interface {
	HasMethod() bool
	Method() hash.Hash
	HasFrom() bool
	From() hash.Hash
	HasTo() bool
	To() hash.Hash
	HasExpireOn() bool
	ExpireOn() time.Time
	HasNotes() bool
	Notes() string
}
