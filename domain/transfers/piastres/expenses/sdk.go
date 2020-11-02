package expenses

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

// Adapter represents a expense adapter
type Adapter interface {
	ToExpense(js []byte) (Expense, error)
	ToJSON(expense Expense) ([]byte, error)
}

// Builder represents an expense builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithAmount(amount uint64) Builder
	From(from []hash.Hash) Builder
	WithLock(lock hash.Hash) Builder
	WithSignatures(signatures []signature.RingSignature) Builder
	WithRemaining(remaining hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Expense, error)
}

// Expense represents an expense
type Expense interface {
	entities.Immutable
	Amount() uint64
	From() []hash.Hash
	Lock() hash.Hash
	Signatures() []signature.RingSignature
	HasRemaining() bool
	Remaining() *hash.Hash
}

// Repository represents an expense repository
type Repository interface {
	Retrieve(hash hash.Hash) (Expense, error)
}

// Service represents the expense service
type Service interface {
	Save(expense Expense) error
}
