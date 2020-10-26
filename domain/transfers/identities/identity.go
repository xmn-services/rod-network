package identities

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type identity struct {
	mutable entities.Mutable
	seed    string
	name    string
	root    string
	wallets []hash.Hash
	buckets []hash.Hash
}

func createIdentityFromJSON(ins *jsonIdentity) (Identity, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	builder := NewBuilder().Create().
		WithHash(*hsh).
		WithSeed(ins.Seed).
		WithName(ins.Name).
		WithRoot(ins.Root).
		CreatedOn(ins.CreatedOn)

	if len(ins.Wallets) > 0 {
		wallets := []hash.Hash{}
		for _, oneWallet := range ins.Wallets {
			hsh, err := hashAdapter.FromString(oneWallet)
			if err != nil {
				return nil, err
			}

			wallets = append(wallets, *hsh)
		}

		builder.WithWallets(wallets)
	}

	if len(ins.Buckets) > 0 {
		buckets := []hash.Hash{}
		for _, oneBucket := range ins.Buckets {
			hsh, err := hashAdapter.FromString(oneBucket)
			if err != nil {
				return nil, err
			}

			buckets = append(buckets, *hsh)
		}

		builder.WithBuckets(buckets)
	}

	return builder.Now()
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
	wallets []hash.Hash,
) Identity {
	return createIdentityInternally(mutable, seed, name, root, wallets, nil)
}

func createIdentityWithBuckets(
	mutable entities.Mutable,
	seed string,
	name string,
	root string,
	buckets []hash.Hash,
) Identity {
	return createIdentityInternally(mutable, seed, name, root, nil, buckets)
}

func createIdentityWithWalletsAndBuckets(
	mutable entities.Mutable,
	seed string,
	name string,
	root string,
	wallets []hash.Hash,
	buckets []hash.Hash,
) Identity {
	return createIdentityInternally(mutable, seed, name, root, wallets, buckets)
}

func createIdentityInternally(
	mutable entities.Mutable,
	seed string,
	name string,
	root string,
	wallets []hash.Hash,
	buckets []hash.Hash,
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
func (obj *identity) Wallets() []hash.Hash {
	return obj.wallets
}

// HasBuckets returns true if there is buckets, false otherwise
func (obj *identity) HasBuckets() bool {
	return obj.buckets != nil
}

// Buckets returns the buckets, if any
func (obj *identity) Buckets() []hash.Hash {
	return obj.buckets
}

// MarshalJSON converts the instance to JSON
func (obj *identity) MarshalJSON() ([]byte, error) {
	ins := createJSONIdentityFromIdentity(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *identity) UnmarshalJSON(data []byte) error {
	ins := new(jsonIdentity)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createIdentityFromJSON(ins)
	if err != nil {
		return err
	}

	insIdentity := pr.(*identity)
	obj.mutable = insIdentity.mutable
	obj.seed = insIdentity.seed
	obj.name = insIdentity.name
	obj.root = insIdentity.root
	obj.wallets = insIdentity.wallets
	obj.buckets = insIdentity.buckets
	return nil
}
