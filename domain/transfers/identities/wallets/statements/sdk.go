package statements

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// Adapter represents a statement adapter
type Adapter interface {
	ToStatement(js []byte) (Statement, error)
	ToJSON(statement Statement) ([]byte, error)
}

// Builder represents the statement builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithIncoming(incoming []hash.Hash) Builder
	WithOutgoing(outgoing []hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Statement, error)
}

// Statement represents a statement
type Statement interface {
	entities.Immutable
	HasIncoming() bool
	Incoming() []hash.Hash
	HasOutgoing() bool
	Outgoing() []hash.Hash
}

// Repository represents a statement repository
type Repository interface {
	Retrieve(hash hash.Hash) (Statement, error)
}

// Service represents a statement service
type Service interface {
	Save(statement Statement) error
	Delete(statement Statement) error
}
