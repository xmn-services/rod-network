package buckets

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/buckets/files"
	transfer_bucket "github.com/xmn-services/rod-network/domain/transfers/buckets"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

// NewService creates a new service instance
func NewService(
	fileService files.Service,
	repository Repository,
	trService transfer_bucket.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, fileService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	fileRepository files.Repository,
	trRepository transfer_bucket.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(fileRepository, trRepository, builder)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	hashTreeBuilder := hashtree.NewBuilder()
	trBuilder := transfer_bucket.NewBuilder()
	return createAdapter(hashTreeBuilder, trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the bucket adapter
type Adapter interface {
	ToTransfer(bucket Bucket) (transfer_bucket.Bucket, error)
}

// Builder represents the bucket builder
type Builder interface {
	Create() Builder
	WithFiles(files []files.File) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Bucket, error)
}

// Bucket represents the bucket
type Bucket interface {
	entities.Immutable
	Files() []files.File
	FileByPath(path string) (files.File, error)
}

// Repository represents a bucket bucket repository
type Repository interface {
	Retrieve(hash hash.Hash) (Bucket, error)
}

// Service represents a bucket bucket service
type Service interface {
	Save(bucket Bucket) error
	Delete(bucket Bucket) error
}
