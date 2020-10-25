package identities

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/buckets"
	"github.com/xmn-services/rod-network/domain/memory/identities/wallets"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type identity struct {
	mutable entities.Mutable
	seed    string
	name    string
	root    string
	wallets []wallets.Wallet
	buckets []buckets.Bucket
}

func createIdentity(
	mutable entities.Mutable,
	seed string,
	name string,
	root string,
) Identity {
	return createIdentityInternally(mutable, seed, name, root, nil, nil)
}

func createIdentityWithWallets(
	mutable entities.Mutable,
	seed string,
	name string,
	root string,
	wallets []wallets.Wallet,
) Identity {
	return createIdentityInternally(mutable, seed, name, root, wallets, nil)
}

func createIdentityWithBuckets(
	mutable entities.Mutable,
	seed string,
	name string,
	root string,
	buckets []buckets.Bucket,
) Identity {
	return createIdentityInternally(mutable, seed, name, root, nil, buckets)
}

func createIdentityWithWalletsAndBuckets(
	mutable entities.Mutable,
	seed string,
	name string,
	root string,
	wallets []wallets.Wallet,
	buckets []buckets.Bucket,
) Identity {
	return createIdentityInternally(mutable, seed, name, root, wallets, buckets)
}

func createIdentityInternally(
	mutable entities.Mutable,
	seed string,
	name string,
	root string,
	wallets []wallets.Wallet,
	buckets []buckets.Bucket,
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

// LastUpdatedOn returns the lastUpdatedOn time
func (obj *identity) LastUpdatedOn() time.Time {
	return obj.mutable.LastUpdatedOn()
}

// CreatedOn returns the creation time
func (obj *identity) CreatedOn() time.Time {
	return obj.mutable.CreatedOn()
}

// HasWallets returns true if there is wallets, false otherwise
func (obj *identity) HasWallets() bool {
	return obj.wallets != nil
}

// Wallets returns the wallets, if any
func (obj *identity) Wallets() []wallets.Wallet {
	return obj.wallets
}

// HasBuckets returns true if there is buckets, false otherwise
func (obj *identity) HasBuckets() bool {
	return obj.buckets != nil
}

// Buckets returns the buckets, if any
func (obj *identity) Buckets() []buckets.Bucket {
	return obj.buckets
}
