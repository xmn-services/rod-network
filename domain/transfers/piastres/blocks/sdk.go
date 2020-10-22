package blocks

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
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

// Adapter represents a block adapter
type Adapter interface {
	ToBlock(js []byte) (Block, error)
	ToJSON(block Block) ([]byte, error)
}

// Builder represents the block builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithAddress(address hash.Hash) Builder
	WithTransactions(trx hashtree.HashTree) Builder
	WithAmount(amount uint) Builder
	WithAdditional(additional uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Block, error)
}

// Block represents a block of transactions
type Block interface {
	entities.Immutable
	Address() hash.Hash
	Transactions() hashtree.HashTree
	Amount() uint
	Additional() uint
}

// Repository represents a block repository
type Repository interface {
	Retrieve(hash hash.Hash) (Block, error)
}

// Service represents the block service
type Service interface {
	Save(block Block) error
}
