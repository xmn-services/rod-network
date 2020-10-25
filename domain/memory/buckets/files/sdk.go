package files

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
	"github.com/xmn-services/rod-network/domain/memory/buckets/files/chunks"
	transfer_file "github.com/xmn-services/rod-network/domain/transfers/buckets/files"
)

// NewService creates a new service instance
func NewService(
	chunksService chunks.Service,
	repository Repository,
	trService transfer_file.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, chunksService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	chunkRepository chunks.Repository,
	trRepository transfer_file.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(chunkRepository, trRepository, builder)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	hashTreeBuilder := hashtree.NewBuilder()
	trBuilder := transfer_file.NewBuilder()
	return createAdapter(hashTreeBuilder, trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the file adapter
type Adapter interface {
	ToTransfer(file File) (transfer_file.File, error)
}

// Builder represents the file builder
type Builder interface {
	Create() Builder
	WithRelativePath(relativePath string) Builder
	WithChunks(chunks []chunks.Chunk) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (File, error)
}

// File represents a file
type File interface {
	entities.Immutable
	RelativePath() string
	Chunks() []chunks.Chunk
	ChunkByHash(hash hash.Hash) (chunks.Chunk, error)
}

// Repository represents a file repository
type Repository interface {
	Retrieve(hash hash.Hash) (File, error)
	RetrieveAll(hashes []hash.Hash) ([]File, error)
}

// Service represents a file service
type Service interface {
	Save(file File) error
	SaveAll(files []File) error
	Delete(file File) error
	DeleteAll(files []File) error
}
