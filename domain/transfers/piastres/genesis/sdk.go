package genesis

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

// Adapter represents a genesis adapter
type Adapter interface {
	ToGenesis(js []byte) (Genesis, error)
	ToJSON(genesis Genesis) ([]byte, error)
}

// Builder represents a genesis builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithBlockDifficultyBase(blockDiffBase uint) Builder
	WithBlockDifficultyIncreasePerTrx(blockDiffIncreasePerTrx float64) Builder
	WithLinkDifficulty(link uint) Builder
	WithBill(bill hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Genesis, error)
}

// Genesis represents the genesis
type Genesis interface {
	entities.Immutable
	Bill() hash.Hash
	BlockDifficultyBase() uint
	BlockDifficultyIncreasePerTrx() float64
	LinkDifficulty() uint
}

// Repository repreents the genesis repository
type Repository interface {
	Retrieve() (Genesis, error)
}

// Service represents the genesis service
type Service interface {
	Save(genesis Genesis) error
}
