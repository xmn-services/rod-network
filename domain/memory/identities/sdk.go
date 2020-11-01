package identities

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/identities/buckets"
	"github.com/xmn-services/rod-network/domain/memory/identities/buckets/bucket"
	"github.com/xmn-services/rod-network/domain/memory/identities/wallets"
	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/wallet"
	"github.com/xmn-services/rod-network/libs/entities"
)

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithSeed(seed string) Builder
	WithName(name string) Builder
	WithRoot(root string) Builder
	WithWallets(wallets []wallet.Wallet) Builder
	WithBuckets(buckets []bucket.Bucket) Builder
	CreatedOn(createdOn time.Time) Builder
	LastUpdatedOn(lastUpdatedOn time.Time) Builder
	Now() (Identity, error)
}

// Identity represents the identity
type Identity interface {
	entities.Mutable
	Seed() string
	Name() string
	Root() string
	Wallets() wallets.Wallets
	Buckets() buckets.Buckets
}

// Repository represents an identity repository
type Repository interface {
	Retrieve(name string, seed string, password string) (Identity, error)
}

// Service represents an identity service
type Service interface {
	Insert(identity Identity, password string) error
	Update(updated Identity, name string, seed string, password string, newPassword string) error
	Delete(name string, seed string, password string) error
}
