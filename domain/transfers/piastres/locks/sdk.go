package locks

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

// Adapter represents a lock adapter
type Adapter interface {
	ToLock(js []byte) (Lock, error)
	ToJSON(lock Lock) ([]byte, error)
}

// Builder represents a lock builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithShareHolders(shareHolders hashtree.HashTree) Builder
	WithTreeshold(treeshold uint) Builder
	WithAmount(amount uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Lock, error)
}

// Lock represents a lock
type Lock interface {
	entities.Immutable
	ShareHolders() hashtree.HashTree
	Treeshold() uint
	Amount() uint
}

// Repository represents a lock repository
type Repository interface {
	Retrieve(hash hash.Hash) (Lock, error)
}

// Service represents the lock service
type Service interface {
	Save(lock Lock) error
}
