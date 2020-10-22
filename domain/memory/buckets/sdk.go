package buckets

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/buckets/informations"
	transfer_bucket "github.com/xmn-services/rod-network/domain/transfers/buckets"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

// NewService creates a new service instance
func NewService(
	informationService informations.Service,
	repository Repository,
	trService transfer_bucket.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, informationService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	informationRepository informations.Repository,
	trRepository transfer_bucket.Repository,
) Repository {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()
	return createRepository(hashAdapter, informationRepository, trRepository, builder)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_bucket.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	pkAdapter := encryption.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, pkAdapter, immutableBuilder)
}

// Adapter returns the bucket adapter
type Adapter interface {
	ToTransfer(bucket Bucket) (transfer_bucket.Bucket, error)
}

// Builder represents a bucket builder
type Builder interface {
	Create() Builder
	WithInformation(information informations.Information) Builder
	WithAbsolutePath(absolutePath string) Builder
	WithPrivateKey(pk encryption.PrivateKey) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Bucket, error)
}

// Bucket represents a bucket
type Bucket interface {
	entities.Immutable
	Information() informations.Information
	AbsolutePath() string
	PrivateKey() encryption.PrivateKey
}

// Repository represents a bucket repository
type Repository interface {
	RetrieveAll() ([]Bucket, error)
	Retrieve(absolutePath string) (Bucket, error)
}

// Service represents a bucket service
type Service interface {
	Save(bucket Bucket) error
	Delete(bucket Bucket) error
}
