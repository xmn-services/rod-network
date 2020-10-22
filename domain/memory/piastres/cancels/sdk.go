package cancels

import (
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	transfer_cancel "github.com/xmn-services/rod-network/domain/transfers/piastres/cancels"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	expenseService expenses.Service,
	lockService locks.Service,
	trService transfer_cancel.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, expenseService, lockService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	expenseRepository expenses.Repository,
	lockRepository locks.Repository,
	trRepository transfer_cancel.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(builder, expenseRepository, lockRepository, trRepository)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_cancel.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the cancel adapter
type Adapter interface {
	ToTransfer(cancel Cancel) (transfer_cancel.Cancel, error)
	ToJSON(ins Cancel) *JSONCancel
	ToCancel(ins *JSONCancel) (Cancel, error)
}

// Builder represents a cancel builder
type Builder interface {
	Create() Builder
	WithExpense(expense expenses.Expense) Builder
	WithLock(lock locks.Lock) Builder
	WithSignatures(signatures []signature.RingSignature) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Cancel, error)
}

// Cancel represents an expense cancel
type Cancel interface {
	entities.Immutable
	Expense() expenses.Expense
	Lock() locks.Lock
	Signatures() []signature.RingSignature
}

// Repository represents a cancel repository
type Repository interface {
	Retrieve(hash hash.Hash) (Cancel, error)
}

// Service represents the cancel service
type Service interface {
	Save(cancel Cancel) error
}
