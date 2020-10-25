package bills

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	transfer_bill "github.com/xmn-services/rod-network/domain/transfers/piastres/bills"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// NewService creates a new service instance
func NewService(repository Repository, lockService locks.Service, trService transfer_bill.Service) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, lockService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(lockRepository locks.Repository, trRepository transfer_bill.Repository) Repository {
	builder := NewBuilder()
	return createRepository(builder, lockRepository, trRepository)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_bill.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the bill adapter
type Adapter interface {
	ToTransfer(bill Bill) (transfer_bill.Bill, error)
	ToJSON(ins Bill) *JSONBill
	ToBill(ins *JSONBill) (Bill, error)
}

// Builder represents the bill builder
type Builder interface {
	Create() Builder
	WithLock(lock locks.Lock) Builder
	WithAmount(amount uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Bill, error)
}

// Bill represent a piastre bill
type Bill interface {
	entities.Immutable
	Lock() locks.Lock
	Amount() uint
}

// Repository represents a bill repository
type Repository interface {
	Retrieve(hash hash.Hash) (Bill, error)
}

// Service represents the bill service
type Service interface {
	Save(bill Bill) error
}
