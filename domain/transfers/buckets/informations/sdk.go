package informations

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

// Adapter represents a information adapter
type Adapter interface {
	ToInformation(js []byte) (Information, error)
	ToJSON(information Information) ([]byte, error)
}

// Builder represents the information builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithFiles(files hashtree.HashTree) Builder
	WithAmount(amount uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Information, error)
}

// Information represents the bucket information
type Information interface {
	entities.Immutable
	Files() hashtree.HashTree
	Amount() uint
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
