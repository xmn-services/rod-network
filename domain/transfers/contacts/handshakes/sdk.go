package handshakes

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Adapter represents an handshake adapter
type Adapter interface {
	ToHandshake(js []byte) (Handshake, error)
	ToJSON(handshake Handshake) ([]byte, error)
}

// Builder represents an handshake builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithIncoming(incoming hash.Hash) Builder
	WithOutgoing(outgoing hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Handshake, error)
}

// Handshake represents a contact handshake
type Handshake interface {
	entities.Immutable
	IsIncoming() bool
	Incoming() hash.Hash
	IsOutgoing() bool
	Outgoing() hash.Hash
}

// Repository represents an handshake repository
type Repository interface {
	Retrieve(hash hash.Hash) (Handshake, error)
}

// Service represents an handshake service
type Service interface {
	Save(handshake Handshake) error
	Delete(handshake Handshake) error
}
