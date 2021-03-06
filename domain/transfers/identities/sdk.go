package identities

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

// NewService creates a new service instance
func NewService(repository Repository, fileService file.Service) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, fileService)
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
	mutableBuilder := entities.NewMutableBuilder()
	return createBuilder(mutableBuilder)
}

// Adapter represents an identity adapter
type Adapter interface {
	ToIdentity(js []byte) (Identity, error)
	ToJSON(identity Identity) ([]byte, error)
}

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithSeed(seed string) Builder
	WithName(name string) Builder
	WithRoot(root string) Builder
	WithWallets(wallets []hash.Hash) Builder
	WithBuckets(buckets []hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	LastUpdatedOn(lastUpdatedOn time.Time) Builder
	Now() (Identity, error)
}

// Identity represents the identity
type Identity interface {
	entities.Mutable
	Seed() string
	Name() string
	Root() string
	HasWallets() bool
	Wallets() []hash.Hash
	HasBuckets() bool
	Buckets() []hash.Hash
}

// Repository represents an identity repository
type Repository interface {
	Retrieve(hash hash.Hash, seed string) (Identity, error)
}

// Service represents an identity service
type Service interface {
	Save(identity Identity) error
	Delete(identity Identity) error
}
