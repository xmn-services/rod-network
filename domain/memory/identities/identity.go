package identities

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/identities/buckets"
	"github.com/xmn-services/rod-network/domain/memory/identities/wallets"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type identity struct {
	mutable entities.Mutable
	seed    string
	name    string
	root    string
	wallets wallets.Wallets
	buckets buckets.Buckets
}

func createIdentity(
	mutable entities.Mutable,
	seed string,
	name string,
	root string,
	wallets wallets.Wallets,
	buckets buckets.Buckets,
) Identity {
	return createIdentityInternally(mutable, seed, name, root, wallets, buckets)
}

func createIdentityInternally(
	mutable entities.Mutable,
	seed string,
	name string,
	root string,
	wallets wallets.Wallets,
	buckets buckets.Buckets,
) Identity {
	out := identity{
		mutable: mutable,
		seed:    seed,
		name:    name,
		root:    root,
		wallets: wallets,
		buckets: buckets,
	}

	return &out
}

// Hash returns the hash
func (obj *identity) Hash() hash.Hash {
	return obj.mutable.Hash()
}

// Seed returns the seed
func (obj *identity) Seed() string {
	return obj.seed
}

// Name returns the name
func (obj *identity) Name() string {
	return obj.name
}

// Root returns the root
func (obj *identity) Root() string {
	return obj.root
}

// Wallets returns the wallets
func (obj *identity) Wallets() wallets.Wallets {
	return obj.wallets
}

// Buckets returns the buckets
func (obj *identity) Buckets() buckets.Buckets {
	return obj.buckets
}

// LastUpdatedOn returns the lastUpdatedOn time
func (obj *identity) LastUpdatedOn() time.Time {
	return obj.mutable.LastUpdatedOn()
}

// CreatedOn returns the creation time
func (obj *identity) CreatedOn() time.Time {
	return obj.mutable.CreatedOn()
}
