package entries

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
	adapter := createAdapter()
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

// Adapter represents an entry adapter
type Adapter interface {
	ToEntry(js []byte) (Entry, error)
	ToJSON(entry Entry) ([]byte, error)
}

// Builder represents an entry builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithTransactions(trx []hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Entry, error)
}

// Entry represents a statement entry
type Entry interface {
	entities.Immutable
	Name() string
	Transactions() []hash.Hash
	HasDescription() bool
	Description() string
}

// Repository represents an entry repository
type Repository interface {
	Retrieve(hash hash.Hash) (Entry, error)
}

// Service represents an entry service
type Service interface {
	Save(entry Entry) error
	Delete(entry Entry) error
}
