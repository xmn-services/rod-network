package locks

import (
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	shareHolders     hashtree.HashTree
	treeshold        uint
	amount           uint
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		shareHolders:     nil,
		treeshold:        0,
		amount:           0,
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

// WithShareHolders adds shareholder's hashes to the builder
func (app *builder) WithShareHolders(shareHolders hashtree.HashTree) Builder {
	app.shareHolders = shareHolders
	return app
}

// WithTreeshold adds a treeshold to the builder
func (app *builder) WithTreeshold(treeshold uint) Builder {
	app.treeshold = treeshold
	return app
}

// WithAmount adds the amount of shareholders to the builder
func (app *builder) WithAmount(amount uint) Builder {
	app.amount = amount
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Lock instance
func (app *builder) Now() (Lock, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Lock instance")
	}

	if app.shareHolders == nil {
		return nil, errors.New("the shareholder hashes are mandatory in order to build a Lock instance")
	}

	if app.treeshold <= 0 {
		return nil, errors.New("the treeshold is mandatory in order to build a Lock instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount of shareholders must be greater than zero (0)")
	}

	leaves := app.shareHolders.Compact().Leaves().Leaves()
	max := len(leaves)
	if app.amount > uint(max) {
		str := fmt.Sprintf("the maximum amount of shareholders is: %d, %d provided", max, app.amount)
		return nil, errors.New(str)
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createLock(immutable, app.shareHolders, app.treeshold, app.amount), nil
}
