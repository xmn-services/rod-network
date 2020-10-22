package genesis

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder        entities.ImmutableBuilder
	hash                    *hash.Hash
	bill                    *hash.Hash
	blockDiffBase           uint
	blockDiffIncreasePerTrx float64
	linkDiff                uint
	createdOn               *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder:        immutableBuilder,
		hash:                    nil,
		bill:                    nil,
		blockDiffBase:           0,
		blockDiffIncreasePerTrx: 0,
		linkDiff:                0,
		createdOn:               nil,
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

// WithBlockDifficultyBase adds a block difficulty base to the builder
func (app *builder) WithBlockDifficultyBase(blockDiffBase uint) Builder {
	app.blockDiffBase = blockDiffBase
	return app
}

// WithBlockDifficultyIncreasePerTrx adds a block difficulty increase per trx to the builder
func (app *builder) WithBlockDifficultyIncreasePerTrx(blockDiffIncreasePerTrx float64) Builder {
	app.blockDiffIncreasePerTrx = blockDiffIncreasePerTrx
	return app
}

// WithLinkDifficulty adds a link difficulty increase per trx to the builder
func (app *builder) WithLinkDifficulty(linkDiff uint) Builder {
	app.linkDiff = linkDiff
	return app
}

// WithBill adds a bill to the builder
func (app *builder) WithBill(bill hash.Hash) Builder {
	app.bill = &bill
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Genesis instance
func (app *builder) Now() (Genesis, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Genesis instance")
	}

	if app.bill == nil {
		return nil, errors.New("the bill hash is mandatory in order to build a Genesis instance")
	}

	if app.blockDiffBase <= 0 {
		return nil, errors.New("the block difficulty base must be greater than zero (0) in order to build a Genesis instance")
	}

	if app.blockDiffIncreasePerTrx <= 0 {
		return nil, errors.New("the block difficulty increasePerTrx must be greater than zero (0) in order to build a Genesis instance")
	}

	if app.linkDiff <= 0 {
		return nil, errors.New("the link difficulty must be greater than zero (0) in order to build a Genesis instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createGenesis(
		immutable,
		*app.bill,
		app.blockDiffBase,
		app.blockDiffIncreasePerTrx,
		app.linkDiff,
	), nil
}
