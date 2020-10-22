package mined

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

// NewService creates a new service instance
func NewService(fileService file.Service) Service {
	adapter := NewAdapter()
	return createService(adapter, fileService)
}

// NewRepository creates a new repository instance
func NewRepository(fileRepository file.Repository) Repository {
	adapter := NewAdapter()
	return createRepository(adapter, fileRepository)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(immutableBuilder)
}

// Adapter represents a mined link adapter
type Adapter interface {
	ToLink(js []byte) (Link, error)
	ToJSON(link Link) ([]byte, error)
}

// Builder represents a link builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithLink(link hash.Hash) Builder
	WithMining(mining string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Link, error)
}

// Link represents a mined link
type Link interface {
	entities.Immutable
	Link() hash.Hash
	Mining() string
}

// Repository represents a link repository
type Repository interface {
	Retrieve(hash hash.Hash) (Link, error)
}

// Service represents the link service
type Service interface {
	Save(link Link) error
}
