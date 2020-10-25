package contacts

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Adapter represents a contact adapter
type Adapter interface {
	ToContact(js []byte) (Contact, error)
	ToJSON(contact Contact) ([]byte, error)
}

// Builder represents a contact builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithHandshake(handshake hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Contact, error)
}

// Contact represents a contact
type Contact interface {
	entities.Immutable
	Name() string
	Description() string
	Handshake() hash.Hash
}

// Repository represents a contact repository
type Repository interface {
	Retrieve(hash hash.Hash) (Contact, error)
}

// Service represents a contact service
type Service interface {
	Save(contact Contact) error
}
