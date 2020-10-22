package chunks

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

// NewService creates a new service instance
func NewService(fileService file.Service) Service {
	adapter := NewAdapter()
	return createService(adapter, fileService)
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
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(immutableBuilder)
}

// Adapter represents a chunk adapter
type Adapter interface {
	ToChunk(js []byte) (Chunk, error)
	ToJSON(chunk Chunk) ([]byte, error)
}

// Builder represents the chunk builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
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
	Delete(chunk Chunk) error
}
