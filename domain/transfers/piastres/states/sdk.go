package states

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

// Adapter represents a state adapter
type Adapter interface {
	ToState(js []byte) (State, error)
	ToJSON(state State) ([]byte, error)
}

// Builder represents a state builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithChain(chain hash.Hash) Builder
	WithPrevious(prev hash.Hash) Builder
	WithHeight(height uint) Builder
	WithTransactions(trx hashtree.HashTree) Builder
	WithAmount(amount uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (State, error)
}

// State represents a chain's state
type State interface {
	entities.Immutable
	Chain() hash.Hash
	Previous() hash.Hash
	Height() uint
	Transactions() hashtree.HashTree
	Amount() uint
}

// Repository represents a state repository
type Repository interface {
	Retrieve(chainHash hash.Hash, height uint) (State, error)
}

// Service represents a state service
type Service interface {
	Save(state State) error
}
