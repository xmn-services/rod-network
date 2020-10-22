package buckets

import (
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/entities"
	libs_file "github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
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

// Adapter represents a bucket adapter
type Adapter interface {
	ToBucket(js []byte) (Bucket, error)
	ToJSON(bucket Bucket) ([]byte, error)
}

// Builder represents a bucket builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithInformation(information hash.Hash) Builder
	WithAbsolutePath(absolutePath string) Builder
	WithPrivateKey(pk encryption.PrivateKey) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Bucket, error)
}

// Bucket represents a bucket
type Bucket interface {
	entities.Immutable
	Information() hash.Hash
	AbsolutePath() string
	PrivateKey() encryption.PrivateKey
}

// Repository represents a bucket repository
type Repository interface {
	RetrieveAll() ([]Bucket, error)
	Retrieve(hash hash.Hash) (Bucket, error)
}

// Service represents a bucket service
type Service interface {
	Save(bucket Bucket) error
	Delete(bucket Bucket) error
}
