package chunks

import (
	"time"

	transfer_chunk "github.com/xmn-services/rod-network/domain/transfers/buckets/files/chunks"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// NewService creates a new service instance
func NewService(repository Repository, trService transfer_chunk.Service) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, trService)
}

// NewRepository creates a new repository instance
func NewRepository(trRepository transfer_chunk.Repository) Repository {
	builder := NewBuilder()
	return createRepository(trRepository, builder)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_chunk.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the chunk adapter
type Adapter interface {
	ToTransfer(chunk Chunk) (transfer_chunk.Chunk, error)
}

// Builder represents the chunk builder
type Builder interface {
	Create() Builder
	WithSizeInBytes(sizeInBytes uint) Builder
	WithData(data hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Chunk, error)
}

// Chunk represents a chunk
type Chunk interface {
	entities.Immutable
	SizeInBytes() uint
	Data() hash.Hash
}

// Repository retrieves the chunk repository
type Repository interface {
	Retrieve(hash hash.Hash) (Chunk, error)
	RetrieveAll(hashes []hash.Hash) ([]Chunk, error)
}

// Service represents the chunk service
type Service interface {
	Save(chunk Chunk) error
	SaveAll(chunks []Chunk) error
	Delete(chunk Chunk) error
	DeleteAll(chunks []Chunk) error
}
