package mined

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/domain/memory/piastres/blocks"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	block            blocks.Block
	mining           string
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		block:            nil,
		mining:           "",
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithBlock adds a block to the builder
func (app *builder) WithBlock(block blocks.Block) Builder {
	app.block = block
	return app
}

// WithMining adds mining results to the builder
func (app *builder) WithMining(mining string) Builder {
	app.mining = mining
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Block instance
func (app *builder) Now() (Block, error) {
	if app.block == nil {
		return nil, errors.New("the block is mandatory in order to build a mined Block instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.block.Hash().Bytes(),
		[]byte(app.mining),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createBlock(immutable, app.block, app.mining), nil
}
