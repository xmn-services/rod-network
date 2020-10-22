package states

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	transfer_state "github.com/xmn-services/rod-network/domain/transfers/piastres/states"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	trService transfer_state.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	trxRepository transactions.Repository,
	trRepository transfer_state.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(trxRepository, trRepository, builder)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	htBuilder := hashtree.NewBuilder()
	trBuilder := transfer_state.NewBuilder()
	return createAdapter(htBuilder, trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter represents a state adapter
type Adapter interface {
	ToTransfer(state State) (transfer_state.State, error)
}

// Builder represents a state builder
type Builder interface {
	Create() Builder
	WithChain(chain hash.Hash) Builder
	WithPrevious(prev hash.Hash) Builder
	WithHeight(height uint) Builder
	WithTransactions(trx []transactions.Transaction) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (State, error)
}

// State represents a chain's state
type State interface {
	entities.Immutable
	Chain() hash.Hash
	Previous() hash.Hash
	Height() uint
	Transactions() []transactions.Transaction
}

// Repository represents a state repository
type Repository interface {
	Retrieve(chain hash.Hash, height uint) (State, error)
}

// Service represents a state service
type Service interface {
	Save(state State) error
}
