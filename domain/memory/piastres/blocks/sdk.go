package blocks

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	"github.com/xmn-services/rod-network/domain/memory/piastres/transactions"
	transfer_block "github.com/xmn-services/rod-network/domain/transfers/piastres/blocks"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	trxService transactions.Service,
	trService transfer_block.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, trxService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	genesisRepository genesis.Repository,
	trxRepository transactions.Repository,
	trRepository transfer_block.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(builder, genesisRepository, trxRepository, trRepository)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	hashTreeBuilder := hashtree.NewBuilder()
	trBuilder := transfer_block.NewBuilder()
	return createAdapter(hashTreeBuilder, trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the block adapter
type Adapter interface {
	ToTransfer(block Block) (transfer_block.Block, error)
	ToJSON(block Block) *JSONBlock
	ToBlock(ins *JSONBlock) (Block, error)
}

// Builder represents the block builder
type Builder interface {
	Create() Builder
	WithAddress(address hash.Hash) Builder
	WithGenesis(gen genesis.Genesis) Builder
	WithAdditional(additional uint) Builder
	WithTransactions(trx []transactions.Transaction) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Block, error)
}

// Block represents a block of transactions
type Block interface {
	entities.Immutable
	Address() hash.Hash
	Genesis() genesis.Genesis
	Additional() uint
	Transactions() []transactions.Transaction
}

// Repository represents a block repository
type Repository interface {
	Retrieve(hash hash.Hash) (Block, error)
}

// Service represents the block service
type Service interface {
	Save(block Block) error
}
