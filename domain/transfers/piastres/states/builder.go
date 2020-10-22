package states

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
	chain            *hash.Hash
	prev             *hash.Hash
	height           uint
	trx              hashtree.HashTree
	amount           uint
	createdOn        *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		chain:            nil,
		prev:             nil,
		height:           0,
		trx:              nil,
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

// WithChain adds a chain hash to the builder
func (app *builder) WithChain(chain hash.Hash) Builder {
	app.chain = &chain
	return app
}

// WithPrevious adds a previous hash to the builder
func (app *builder) WithPrevious(prev hash.Hash) Builder {
	app.prev = &prev
	return app
}

// WithHeight adds an height to the builder
func (app *builder) WithHeight(height uint) Builder {
	app.height = height
	return app
}

// WithTransactions add transactions to the builder
func (app *builder) WithTransactions(trx hashtree.HashTree) Builder {
	app.trx = trx
	return app
}

// WithAmount adds an amount to the builder
func (app *builder) WithAmount(amount uint) Builder {
	app.amount = amount
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new State instance
func (app *builder) Now() (State, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a State instance")
	}

	if app.prev == nil {
		return nil, errors.New("the previous hash is mandatory in order to build a State instance")
	}

	if app.height <= 0 {
		return nil, errors.New("the height must be greater than zero (0) in order to build a State instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transaction's hashtree is mandatory in order to build a State instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount is mandatory in order to build a State instance")
	}

	if app.createdOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a State instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createState(immutable, *app.chain, *app.prev, app.height, app.trx, app.amount), nil
}
