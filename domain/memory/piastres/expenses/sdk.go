package expenses

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks"
	transfer_expense "github.com/xmn-services/rod-network/domain/transfers/piastres/expenses"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	billService bills.Service,
	lockService locks.Service,
	trService transfer_expense.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, billService, lockService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	billRepository bills.Repository,
	lockRepository locks.Repository,
	trRepository transfer_expense.Repository,
) Repository {
	builder := NewBuilder()
	contentBuilder := NewContentBuilder()
	return createRepository(builder, contentBuilder, billRepository, lockRepository, trRepository)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_expense.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// NewContentBuilder creates a new content builder instance
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createContentBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the expense adapter
type Adapter interface {
	ToTransfer(expense Expense) (transfer_expense.Expense, error)
	ToJSON(ins Expense) *JSONExpense
	ToExpense(ins *JSONExpense) (Expense, error)
}

// Builder represents an expense builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignatures(sigs []signature.RingSignature) Builder
	Now() (Expense, error)
}

// Expense represents an expense
type Expense interface {
	entities.Immutable
	Content() Content
	Signatures() []signature.RingSignature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithAmount(amount uint64) ContentBuilder
	From(from []bills.Bill) ContentBuilder
	WithLock(lock locks.Lock) ContentBuilder
	WithRemaining(remaining locks.Lock) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents an expense content
type Content interface {
	entities.Immutable
	Amount() uint64
	From() []bills.Bill
	Lock() locks.Lock
	HasRemaining() bool
	Remaining() locks.Lock
}

// Repository represents an expense repository
type Repository interface {
	Retrieve(hash hash.Hash) (Expense, error)
	RetrieveAll(hashes []hash.Hash) ([]Expense, error)
}

// Service represents the expense service
type Service interface {
	Save(expense Expense) error
	SaveAll(expenses []Expense) error
}
