package contacts

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/contacts/handshakes"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Builder represents a contact builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithHandshake(handshake handshakes.Handshake) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Contact, error)
}

// Contact represents a contact
type Contact interface {
	entities.Immutable
	Name() string
	Description() string
	Handshake() handshakes.Handshake
}

// Repository represents a contact repository
type Repository interface {
	Retrieve(hash hash.Hash) (Contact, error)
}

// Service represents a contact service
type Service interface {
	Save(contact Contact) error
}
