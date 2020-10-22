package genesis

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/bills"
	transfer_genesis "github.com/xmn-services/rod-network/domain/transfers/piastres/genesis"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	billService bills.Service,
	trService transfer_genesis.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, billService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	billRepository bills.Repository,
	trRepository transfer_genesis.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(builder, billRepository, trRepository)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_genesis.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the genesis adapter
type Adapter interface {
	ToTransfer(genesis Genesis) (transfer_genesis.Genesis, error)
	ToJSON(genesis Genesis) *JSONGenesis
	ToGenesis(ins *JSONGenesis) (Genesis, error)
}

// Builder represents a genesis builder
type Builder interface {
	Create() Builder
	WithBlockDifficultyBase(blockDiffBase uint) Builder
	WithBlockDifficultyIncreasePerTrx(blockDiffIncreasePerTrx float64) Builder
	WithLinkDifficulty(link uint) Builder
	WithBill(bill bills.Bill) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Genesis, error)
}

// Genesis represents the genesis
type Genesis interface {
	entities.Immutable
	Bill() bills.Bill
	Difficulty() Difficulty
}

// Difficulty represents the genesis difficulty
type Difficulty interface {
	Block() Block
	Link() uint
}

// Block represents the block difficulty related data
type Block interface {
	Base() uint
	IncreasePerTrx() float64
}

// Repository repreents the genesis repository
type Repository interface {
	Retrieve() (Genesis, error)
}

// Service represents the genesis service
type Service interface {
	Save(genesis Genesis) error
}
