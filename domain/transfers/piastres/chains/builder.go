package chains

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	genesis          *hash.Hash
	root             *hash.Hash
	head             *hash.Hash
	height           uint
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		genesis:          nil,
		root:             nil,
		head:             nil,
		height:           0,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.immutableBuilder)
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.hash = &hash
	return app
}

// WithGenesis adds a genesis hash to the builder
func (app *builder) WithGenesis(gen hash.Hash) Builder {
	app.genesis = &gen
	return app
}

// WithRoot adds a root hash to the builder
func (app *builder) WithRoot(root hash.Hash) Builder {
	app.root = &root
	return app
}

// WithHead adds an head hash to the builder
func (app *builder) WithHead(head hash.Hash) Builder {
	app.head = &head
	return app
}

// WithHeight adds an height to the builder
func (app *builder) WithHeight(height uint) Builder {
	app.height = height
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Chain instance
func (app *builder) Now() (Chain, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Chain instance")
	}

	if app.height <= 0 {
		return nil, errors.New("the height must be greater than zero (0) in order to build a Chain instance")
	}

	if app.genesis == nil {
		return nil, errors.New("the genesis hash is mandatory in order to build a Chain instance")
	}

	if app.root == nil {
		return nil, errors.New("the root hash is mandatory in order to build a Chain instance")
	}

	if app.head == nil {
		return nil, errors.New("the head hash is mandatory in order to build a Chain instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createChain(immutable, *app.genesis, *app.root, *app.head, app.height), nil
}
