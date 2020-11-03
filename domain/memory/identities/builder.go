package identities

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/domain/memory/identities/buckets"
	"github.com/xmn-services/rod-network/domain/memory/identities/buckets/bucket"
	"github.com/xmn-services/rod-network/domain/memory/identities/wallets"
	"github.com/xmn-services/rod-network/domain/memory/identities/wallets/wallet"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter    hash.Adapter
	mutableBuilder entities.MutableBuilder
	walletsBuilder wallets.Builder
	bucketsBuilder buckets.Builder
	seed           string
	name           string
	root           string
	wallets        []wallet.Wallet
	buckets        []bucket.Bucket
	createdOn      *time.Time
	lastUpdatedOn  *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	mutableBuilder entities.MutableBuilder,
	walletsBuilder wallets.Builder,
	bucketsBuilder buckets.Builder,
) Builder {
	out := builder{
		hashAdapter:    hashAdapter,
		mutableBuilder: mutableBuilder,
		walletsBuilder: walletsBuilder,
		bucketsBuilder: bucketsBuilder,
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
	return createBuilder(
		app.hashAdapter,
		app.mutableBuilder,
		app.walletsBuilder,
		app.bucketsBuilder,
	)
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
func (app *builder) WithWallets(wallets []wallet.Wallet) Builder {
	app.wallets = wallets
	return app
}

// WithBuckets add buckets to the builder
func (app *builder) WithBuckets(buckets []bucket.Bucket) Builder {
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
	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build an Identity instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Identity instance")
	}

	if app.root == "" {
		return nil, errors.New("the root is mandatory in order to build an Identity instance")
	}

	data := [][]byte{
		[]byte(app.seed),
		[]byte(app.name),
		[]byte(app.root),
	}

	if app.wallets != nil {
		for _, oneWallet := range app.wallets {
			data = append(data, oneWallet.Hash().Bytes())
		}
	}

	if app.buckets != nil {
		for _, oneBucket := range app.buckets {
			data = append(data, oneBucket.Hash().Bytes())
		}
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	mutable, err := app.mutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).LastUpdatedOn(app.lastUpdatedOn).Now()
	if err != nil {
		return nil, err
	}

	walletsBuilder := app.walletsBuilder.Create()
	if app.wallets != nil && len(app.wallets) > 0 {
		walletsBuilder.WithWallets(app.wallets)
	}

	wallets, err := walletsBuilder.Now()
	if err != nil {
		return nil, err
	}

	bucketsBuilder := app.bucketsBuilder.Create()
	if app.buckets != nil && len(app.buckets) > 0 {
		bucketsBuilder.WithBuckets(app.buckets)
	}

	buckets, err := bucketsBuilder.Now()
	if err != nil {
		return nil, err
	}

	return createIdentity(mutable, app.seed, app.name, app.root, wallets, buckets), nil
}
