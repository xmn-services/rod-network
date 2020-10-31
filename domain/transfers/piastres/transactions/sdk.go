package transactions

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

// Adapter represents a transaction adapter
type Adapter interface {
	ToTransaction(js []byte) (Transaction, error)
	ToJSON(transaction Transaction) ([]byte, error)
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithFees(fees []hash.Hash) Builder
	WithBucket(bucket hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	entities.Immutable
	HasBucket() bool
	Bucket() *hash.Hash
	HasFees() bool
	Fees() []hash.Hash
}

// Repository represents a transaction repository
type Repository interface {
	Retrieve(hash hash.Hash) (Transaction, error)
}

// Service represents the transaction service
type Service interface {
	Save(trx Transaction) error
}
