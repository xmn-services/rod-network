package cancels

import (
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
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

// Adapter represents a cancel adapter
type Adapter interface {
	ToCancel(js []byte) (Cancel, error)
	ToJSON(cancel Cancel) ([]byte, error)
}

// Builder represents a cancel builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithExpense(expense hash.Hash) Builder
	WithLock(lock hash.Hash) Builder
	WithSignatures(signatures []signature.RingSignature) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Cancel, error)
}

// Cancel represents an expense cancel
type Cancel interface {
	entities.Immutable
	Expense() hash.Hash
	Lock() hash.Hash
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
