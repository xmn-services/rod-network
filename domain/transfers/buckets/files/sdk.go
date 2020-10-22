package files

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	libs_file "github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

// NewService creates a new service instance
func NewService(fileService libs_file.Service) Service {
	adapter := NewAdapter()
	return createService(adapter, fileService)
}

// NewRepository creates a new repository instance
func NewRepository(fileRepository libs_file.Repository) Repository {
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

// Adapter represents a file adapter
type Adapter interface {
	ToFile(js []byte) (File, error)
	ToJSON(file File) ([]byte, error)
}

// Builder represents the file builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithRelativePath(relativePath string) Builder
	WithChunks(chunks hashtree.HashTree) Builder
	WithAmount(amount uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (File, error)
}

// File represents a file
type File interface {
	entities.Immutable
	RelativePath() string
	Chunks() hashtree.HashTree
	Amount() uint
}

// Repository represents a file repository
type Repository interface {
	Retrieve(hash hash.Hash) (File, error)
}

// Service represents a file service
type Service interface {
	Save(file File) error
	Delete(file File) error
}
