package mined

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/links"
	transfer_mined_link "github.com/xmn-services/rod-network/domain/transfers/piastres/links/mined"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	linkService links.Service,
	trService transfer_mined_link.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, linkService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	linkRepository links.Repository,
	trRepository transfer_mined_link.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(builder, linkRepository, trRepository)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_mined_link.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the mined link adapter
type Adapter interface {
	ToTransfer(link Link) (transfer_mined_link.Link, error)
	Decode(encoded string) (Link, error)
	Encode(link Link) (string, error)
}

// Builder represents a link builder
type Builder interface {
	Create() Builder
	WithLink(link links.Link) Builder
	WithMining(mining string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Link, error)
}

// Link represents a mined link
type Link interface {
	entities.Immutable
	Link() links.Link
	Mining() string
}

// Repository represents a link repository
type Repository interface {
	Retrieve(hash hash.Hash) (Link, error)
}

// Service represents the link service
type Service interface {
	Save(link Link) error
}
