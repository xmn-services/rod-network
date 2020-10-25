package identities

import (
	"hash"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
)

// Adapter represents an identity adapter
type Adapter interface {
	ToIdentity(js []byte) (Identity, error)
	ToJSON(identity Identity) ([]byte, error)
}

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithSeed(seed string) Builder
	WithName(name string) Builder
	WithRoot(root string) Builder
	WithWallets(wallets []hash.Hash) Builder
	WithBuckets(buckets []hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Identity, error)
}

// Identity represents the identity
type Identity interface {
	entities.Mutable
	Seed() string
	Name() string
	Root() string
	HasWallets() bool
	Wallets() []hash.Hash
	HasBuckets() bool
	Buckets() []hash.Hash
}

// Repository represents an identity repository
type Repository interface {
	Retrieve(hash hash.Hash) (Identity, error)
}

// Service represents an identity service
type Service interface {
	Save(identity Identity) error
	Update(original Identity, updated Identity) error
	Delete(identity Identity) error
}
