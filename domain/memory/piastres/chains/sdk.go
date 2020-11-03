package chains

import (
	"time"

	mined_block "github.com/xmn-services/rod-network/domain/memory/piastres/blocks/mined"
	"github.com/xmn-services/rod-network/domain/memory/piastres/genesis"
	mined_link "github.com/xmn-services/rod-network/domain/memory/piastres/links/mined"
	transfer_chains "github.com/xmn-services/rod-network/domain/transfers/piastres/chains"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	genesisService genesis.Service,
	blockService mined_block.Service,
	linkService mined_link.Service,
	trService transfer_chains.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, genesisService, blockService, linkService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	genesisRepository genesis.Repository,
	blockRepository mined_block.Repository,
	linkRepository mined_link.Repository,
	trRepository transfer_chains.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(genesisRepository, blockRepository, linkRepository, trRepository, builder)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_chains.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the chain adapter
type Adapter interface {
	ToTransfer(chain Chain) (transfer_chains.Chain, error)
}

// Builder represents the chain builder
type Builder interface {
	Create() Builder
	WithGenesis(gen genesis.Genesis) Builder
	WithRoot(root mined_block.Block) Builder
	WithHead(head mined_link.Link) Builder
	WithTotal(total uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Chain, error)
}

// Chain represents a chain
type Chain interface {
	entities.Immutable
	Genesis() genesis.Genesis
	Root() mined_block.Block
	Head() mined_link.Link
	Total() uint
	Height() uint
}

// Repository represents the chain repository
type Repository interface {
	Retrieve() (Chain, error)
}

// Service represents the chain service
type Service interface {
	Insert(chain Chain) error
	Update(original Chain, updated Chain) error
}
