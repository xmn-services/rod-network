package mined

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
	transfer_block_mined "github.com/xmn-services/rod-network/domain/transfers/piastres/blocks/mined"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	blockService blocks.Service,
	trService transfer_block_mined.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, blockService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	blockRepository blocks.Repository,
	trRepository transfer_block_mined.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(builder, blockRepository, trRepository)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_block_mined.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the block adapter
type Adapter interface {
	ToTransfer(block Block) (transfer_block_mined.Block, error)
}

// Builder represents a mined block builder
type Builder interface {
	Create() Builder
	WithBlock(block blocks.Block) Builder
	WithMining(mining string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Block, error)
}

// Block represents a mined block
type Block interface {
	entities.Immutable
	Block() blocks.Block
	Mining() string
}

// Repository represents a block repository
type Repository interface {
	Retrieve(hash hash.Hash) (Block, error)
}

// Service represents the block service
type Service interface {
	Save(block Block) error
}
