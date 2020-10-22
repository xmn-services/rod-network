package bills

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

// Adapter represents a bill adapter
type Adapter interface {
	ToBill(js []byte) (Bill, error)
	ToJSON(bill Bill) ([]byte, error)
}

// Builder represents the bill builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithLock(lock hash.Hash) Builder
	WithAmount(amount uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Bill, error)
}

// Bill represent a piastre bill
type Bill interface {
	entities.Immutable
	Lock() hash.Hash
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
