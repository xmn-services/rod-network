package locks

import (
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
	"github.com/xmn-services/rod-network/domain/memory/piastres/locks/shareholders"
	transfer_lock "github.com/xmn-services/rod-network/domain/transfers/piastres/locks"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	shareHolderService shareholders.Service,
	trService transfer_lock.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, shareHolderService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	shareHolderRepository shareholders.Repository,
	trRepository transfer_lock.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(shareHolderRepository, trRepository, builder)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	htBuilder := hashtree.NewBuilder()
	trBuilder := transfer_lock.NewBuilder()
	return createAdapter(htBuilder, trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the lock adapter
type Adapter interface {
	ToTransfer(lock Lock) (transfer_lock.Lock, error)
	ToJSON(ins Lock) *JSONLock
	ToLock(ins *JSONLock) (Lock, error)
}

// Builder represents a lock builder
type Builder interface {
	Create() Builder
	WithShareHolders(shareHolders []shareholders.ShareHolder) Builder
	WithTreeshold(treeshold uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Lock, error)
}

// Lock represents a lock
type Lock interface {
	entities.Immutable
	ShareHolders() []shareholders.ShareHolder
	Treeshold() uint
	Unlock(signatures []signature.RingSignature) error
}

// Repository represents a lock repository
type Repository interface {
	Retrieve(hash hash.Hash) (Lock, error)
}

// Service represents the lock service
type Service interface {
	Save(lock Lock) error
}
