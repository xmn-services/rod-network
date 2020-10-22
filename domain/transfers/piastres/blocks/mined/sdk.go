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

// Adapter represents a mined block adapter
type Adapter interface {
	ToBlock(js []byte) (Block, error)
	ToJSON(block Block) ([]byte, error)
}

// Builder represents a mined block builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithBlock(block hash.Hash) Builder
	WithMining(mining string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Block, error)
}

// Block represents a mined block
type Block interface {
	entities.Immutable
	Block() hash.Hash
	Mining() string
}

// Repository represents a block repository
type Repository interface {
	Retrieve(hash hash.Hash) (Block, error)
}

// Service represents the block service
type Service interface {
	Save(block Block) error
}
