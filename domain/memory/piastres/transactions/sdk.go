package transactions

import (
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/cancels"
	"github.com/xmn-services/rod-network/domain/memory/piastres/expenses"
	transfer_transaction "github.com/xmn-services/rod-network/domain/transfers/piastres/transactions"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	expenseService expenses.Service,
	cancelService cancels.Service,
	trService transfer_transaction.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, expenseService, cancelService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	builder Builder,
	expenseRepository expenses.Repository,
	cancelRepository cancels.Repository,
	trRepository transfer_transaction.Repository,
) Repository {
	contentBuilder := NewContentBuilder()
	return createRepository(builder, contentBuilder, expenseRepository, cancelRepository, trRepository)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_transaction.NewBuilder()
	return createAdapter(trBuilder)
}

// NewContentBuilder returns a new content builder instance
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createContentBuilder(hashAdapter, immutableBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder(amountRingKeys uint) Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	pkFactory := signature.NewPrivateKeyFactory()
	return createBuilder(hashAdapter, immutableBuilder, pkFactory, amountRingKeys)
}

// Adapter returns the transaction adapter
type Adapter interface {
	ToTransfer(trx Transaction) (transfer_transaction.Transaction, error)
	ToJSON(trx Transaction) *JSONTransaction
	ToTransaction(ins *JSONTransaction) (Transaction, error)
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignature(signature signature.RingSignature) Builder
	WithPrivateKey(pk signature.PrivateKey) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	entities.Immutable
	Content() Content
	Signature() signature.RingSignature
}

// ContentBuilder represents a content builder instance
type ContentBuilder interface {
	Create() ContentBuilder
	TriggersOn(triggersOn time.Time) ContentBuilder
	ExecutesOnTrigger() ContentBuilder
	WithFees(fees expenses.Expense) ContentBuilder
	WithExpense(expense expenses.Expense) ContentBuilder
	WithCancel(cancel cancels.Cancel) ContentBuilder
	Now() (Content, error)
}

// Content represents the content of a transaction
type Content interface {
	Hash() hash.Hash
	TriggersOn() time.Time
	ExecutesOnTrigger() bool
	IsExpense() bool
	Expense() expenses.Expense
	IsCancel() bool
	Cancel() cancels.Cancel
	HasFees() bool
	Fees() expenses.Expense
}

// Repository represents a transaction repository
type Repository interface {
	Retrieve(hash hash.Hash) (Transaction, error)
	RetrieveAll(hashes []hash.Hash) ([]Transaction, error)
}

// Service represents the transaction service
type Service interface {
	Save(trx Transaction) error
	SaveAll(trx []Transaction) error
}
