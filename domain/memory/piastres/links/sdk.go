package links

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
	transfer_link "github.com/xmn-services/rod-network/domain/transfers/piastres/links"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	blockService blocks.Service,
	trService transfer_link.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, blockService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	blockRepository blocks.Repository,
	trRepository transfer_link.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(builder, blockRepository, trRepository)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_link.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the link adapter
type Adapter interface {
	ToTransfer(link Link) (transfer_link.Link, error)
	ToJSON(link Link) *JSONLink
	ToLink(ins *JSONLink) (Link, error)
}

// Builder represents a link builder
type Builder interface {
	Create() Builder
	WithPreviousLink(prevLink hash.Hash) Builder
	WithNext(next blocks.Block) Builder
	WithIndex(index uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Link, error)
}

// Link represents a block link
type Link interface {
	entities.Immutable
	PreviousLink() hash.Hash
	Next() blocks.Block
	Index() uint
}

// Repository represents a link repository
type Repository interface {
	Retrieve(hash hash.Hash) (Link, error)
}

// Service represents the link service
type Service interface {
	Save(link Link) error
}
