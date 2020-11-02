package addresses

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

// NewService creates a new service instance
func NewService(
	fileService file.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, fileService)
}

// NewRepository creates a new repository instance
func NewRepository(
	fileRepository file.Repository,
) Repository {
	adapter := NewAdapter()
	return createRepository(adapter, fileRepository)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter represents an address adapter
type Adapter interface {
	ToAddress(js *JSONAddress) (Address, error)
	ToJSON(address Address) *JSONAddress
}

// Builder represents an address builder
type Builder interface {
	Create() Builder
	WithSender(sender hash.Hash) Builder
	WithRecipients(recipients []hash.Hash) Builder
	WithSubject(subject hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Address, error)
}

// Address represents a transaction address
type Address interface {
	entities.Immutable
	HasSender() bool
	Sender() *hash.Hash
	HasRecipients() bool
	Recipients() []hash.Hash
	HasSubject() bool
	Subject() *hash.Hash
}

// Repository represents an address repository
type Repository interface {
	Retrieve(hash hash.Hash) (Address, error)
}

// Service represents an address service
type Service interface {
	Save(address Address) error
}
