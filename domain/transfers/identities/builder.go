package identities

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	mutableBuilder entities.MutableBuilder
	hash           *hash.Hash
	seed           string
	name           string
	root           string
	wallets        []hash.Hash
	buckets        []hash.Hash
	createdOn      *time.Time
	lastUpdatedOn  *time.Time
}

func createBuilder(
	mutableBuilder entities.MutableBuilder,
) Builder {
	out := builder{
		mutableBuilder: mutableBuilder,
		hash:           nil,
		seed:           "",
		name:           "",
		root:           "",
		wallets:        nil,
		buckets:        nil,
		createdOn:      nil,
		lastUpdatedOn:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.mutableBuilder)
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.hash = &hash
	return app
}

// WithSeed adds a seed to the builder
func (app *builder) WithSeed(seed string) Builder {
	app.seed = seed
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root string) Builder {
	app.root = root
	return app
}

// WithWallets add wallets to the builder
func (app *builder) WithWallets(wallets []hash.Hash) Builder {
	app.wallets = wallets
	return app
}

// WithBuckets add buckets to the builder
func (app *builder) WithBuckets(buckets []hash.Hash) Builder {
	app.buckets = buckets
	return app
}

// CreatedOn add a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// LastUpdatedOn add a lastUpdatedOn time to the builder
func (app *builder) LastUpdatedOn(lastUpdatedOn time.Time) Builder {
	app.lastUpdatedOn = &lastUpdatedOn
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build an Identity instance")
	}

	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build an Identity instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Identity instance")
	}

	if app.root == "" {
		return nil, errors.New("the root is mandatory in order to build an Identity instance")
	}

	mutable, err := app.mutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).LastUpdatedOn(app.lastUpdatedOn).Now()
	if err != nil {
		return nil, err
	}

	if app.wallets != nil && app.buckets != nil {
		return createIdentityWithWalletsAndBuckets(mutable, app.seed, app.name, app.root, app.wallets, app.buckets), nil
	}

	if app.wallets != nil {
		return createIdentityWithWallets(mutable, app.seed, app.name, app.root, app.wallets), nil
	}

	if app.buckets != nil {
		return createIdentityWithBuckets(mutable, app.seed, app.name, app.root, app.buckets), nil
	}

	return createIdentity(mutable, app.seed, app.name, app.root), nil
}
