package shareholders

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

// Adapter represents a shareholder adapter
type Adapter interface {
	ToShareHolder(js []byte) (ShareHolder, error)
	ToJSON(shareHolder ShareHolder) ([]byte, error)
}

// Builder represents a shareholder builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithKey(key hash.Hash) Builder
	WithPower(power uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (ShareHolder, error)
}

// ShareHolder represents a lock shareholder
type ShareHolder interface {
	entities.Immutable
	Key() hash.Hash
	Power() uint
}

// Repository represents a shareholder repository
type Repository interface {
	Retrieve(hash hash.Hash) (ShareHolder, error)
}

// Service represents the shareHolder service
type Service interface {
	Save(shareHolder ShareHolder) error
}
