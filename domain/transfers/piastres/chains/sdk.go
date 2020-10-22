package chains

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

// NewService creates a new service instance
func NewService(fileService file.Service, fileNameWithExt string) Service {
	adapter := NewAdapter()
	return createService(adapter, fileService, fileNameWithExt)
}

// NewRepository creates a new repository instance
func NewRepository(fileRepository file.Repository, fileNameWithExt string) Repository {
	adapter := NewAdapter()
	return createRepository(adapter, fileRepository, fileNameWithExt)
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

// Adapter represents a chain adapter
type Adapter interface {
	ToChain(js []byte) (Chain, error)
	ToJSON(chain Chain) ([]byte, error)
}

// Builder represents the chain builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithGenesis(gen hash.Hash) Builder
	WithRoot(root hash.Hash) Builder
	WithHead(head hash.Hash) Builder
	WithHeight(height uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Chain, error)
}

// Chain represents a chain
type Chain interface {
	entities.Immutable
	Genesis() hash.Hash
	Root() hash.Hash
	Head() hash.Hash
	Height() uint
}

// Repository represents the chain repository
type Repository interface {
	Retrieve() (Chain, error)
}

// Service represents the chain service
type Service interface {
	Save(chain Chain) error
}
