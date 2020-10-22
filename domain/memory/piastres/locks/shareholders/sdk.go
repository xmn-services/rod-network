package shareholders

import (
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	transfer_lock_shareholder "github.com/xmn-services/rod-network/domain/transfers/piastres/locks/shareholders"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	trService transfer_lock_shareholder.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	trRepository transfer_lock_shareholder.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(trRepository, builder)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_lock_shareholder.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the shareHolder adapter
type Adapter interface {
	ToTransfer(holder ShareHolder) (transfer_lock_shareholder.ShareHolder, error)
	ToJSON(holder ShareHolder) *JSONShareHolder
	ToShareHolder(ins *JSONShareHolder) (transfer_lock_shareholder.ShareHolder, error)
}

// Builder represents a shareholder builder
type Builder interface {
	Create() Builder
	WithKey(key hash.Hash) Builder
	WithPower(power uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (ShareHolder, error)
}

// ShareHolder represents a lock shareholder
type ShareHolder interface {
	entities.Immutable
	Key() hash.Hash
	Power() uint
}

// Repository represents a shareholder repository
type Repository interface {
	Retrieve(hash hash.Hash) (ShareHolder, error)
	RetrieveAll(hashes []hash.Hash) ([]ShareHolder, error)
}

// Service represents the shareHolder service
type Service interface {
	Save(shareHolder ShareHolder) error
	SaveAll(shareHolders []ShareHolder) error
}
