package informations

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
	"github.com/xmn-services/rod-network/domain/memory/buckets/files"
	transfer_information "github.com/xmn-services/rod-network/domain/transfers/buckets/informations"
)

// NewService creates a new service instance
func NewService(
	fileService files.Service,
	repository Repository,
	trService transfer_information.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, fileService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	fileRepository files.Repository,
	trRepository transfer_information.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(fileRepository, trRepository, builder)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	hashTreeBuilder := hashtree.NewBuilder()
	trBuilder := transfer_information.NewBuilder()
	return createAdapter(hashTreeBuilder, trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the information adapter
type Adapter interface {
	ToTransfer(information Information) (transfer_information.Information, error)
}

// Builder represents the information builder
type Builder interface {
	Create() Builder
	WithFiles(files []files.File) Builder
	WithParent(parent Information) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Information, error)
}

// Information represents the bucket information
type Information interface {
	entities.Immutable
	Files() []files.File
	FileByPath(path string) (files.File, error)
	HasParent() bool
	Parent() Information
}

// Repository represents a bucket information repository
type Repository interface {
	Retrieve(hash hash.Hash) (Information, error)
}

// Service represents a bucket information service
type Service interface {
	Save(information Information) error
	Delete(information Information) error
}
