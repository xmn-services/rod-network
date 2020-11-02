package blocks

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
	"github.com/xmn-services/rod-network/libs/hashtree"
)

type builder struct {
	immutableBuilder entities.ImmutableBuilder
	hash             *hash.Hash
	trx              hashtree.HashTree
	amount           uint
	additional       uint
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		trx:              nil,
		amount:           0,
		additional:       0,
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

// WithTransactions add transaction hashtree to the builder
func (app *builder) WithTransactions(trx hashtree.HashTree) Builder {
	app.trx = trx
	return app
}

// WithAmount adds an amount to the builder
func (app *builder) WithAmount(amount uint) Builder {
	app.amount = amount
	return app
}

// WithAdditional adds an additional to the builder
func (app *builder) WithAdditional(additional uint) Builder {
	app.additional = additional
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
		return nil, errors.New("the hash is mandatory in order to build a Block instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transaction hashtree is mandatory in order to build a Block instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount is mandatory in order to build a Block instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createBlock(immutable, app.trx, app.amount, app.additional), nil
}
