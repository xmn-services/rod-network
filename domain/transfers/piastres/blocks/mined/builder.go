package mined

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	block            *hash.Hash
	mining           string
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		block:            nil,
		mining:           "",
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

// WithBlock adds a block hash to the builder
func (app *builder) WithBlock(block hash.Hash) Builder {
	app.block = &block
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
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a mined Block instance")
	}

	if app.block == nil {
		return nil, errors.New("the block is mandatory in order to build a mined Block instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createBlock(immutable, *app.block, app.mining), nil
}
